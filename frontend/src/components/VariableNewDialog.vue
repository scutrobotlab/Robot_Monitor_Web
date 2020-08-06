<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="600px">
      <v-card>
        <v-toolbar dense color="primary">
          <v-toolbar-title>添加变量</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon dark @click="dialog = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-toolbar>
        <v-card-text>
          <v-container>
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-row>
                <v-col cols="12" sm="6" md="6">
                  <v-select
                    :items="[1, 2, 3]"
                    label="板子代号"
                    hint="保持默认为1即可"
                    :rules="[(v) => !!v || '板子代号是必要的']"
                    required
                    v-model="Board"
                  ></v-select>
                </v-col>
                <v-col cols="12" sm="6" md="6">
                  <v-select
                    :items="types"
                    label="变量类型"
                    :rules="[(v) => !!v || '变量类型是必要的']"
                    required
                    v-model="Type"
                  ></v-select>
                </v-col>
                <v-col cols="12" sm="6" md="6">
                  <v-text-field
                    label="变量名"
                    type="text"
                    :rules="[(v) => !!v || '变量名是必要的']"
                    required
                    v-model="Name"
                  ></v-text-field>
                </v-col>
                <v-col cols="12" sm="6" md="6">
                  <v-text-field
                    label="变量地址"
                    type="text"
                    :rules="AddrRules"
                    hint="形如2000ab78"
                    required
                    v-model="Addr"
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" @click="$refs.form.reset()">清空</v-btn>
          <v-btn color="primary" :disabled="!valid" @click="addVariable()">添加</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import axios from "axios";
export default {
  props: ["opt"],
  data: () => ({
    dialog: false,
    valid: true,
    types: [],
    Board: 1,
    Name: "",
    Type: "",
    Addr: "",
    AddrRules: [
      (v) => !!v || "变量地址是必要的",
      (v) => /2[0-9a-f]{7}/.test(v) || "格式错误，应形如2000ab78",
    ],
  }),
  mounted() {
    axios.get("/variable/types").then((response) => {
      this.types = response.data.Types;
    });
  },
  methods: {
    openDialog() {
      this.dialog = true;
    },
    addVariable() {
      if (this.$refs.form.validate()) {
        axios
          .post("/variable-" + this.opt + "/add", {
            Board: 1,
            Name: this.Name,
            Type: this.Type,
            Addr: parseInt(this.Addr, 16),
          })
          .then((response) => {
            if (response.data.status == 0) {
              this.dialog = false;
              this.$toasted.show("变量添加成功");
              this.$emit("getVariables");
            } else if (response.data.status == 22) {
              this.$toasted.error("变量操作时串口错误");
            } else if (response.data.status == 23) {
              this.$toasted.error("重复添加变量");
            }
          });
      }
    },
  },
};
</script>
