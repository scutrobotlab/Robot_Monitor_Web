<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>变量·{{ showtext }}</v-list-item-title>
      <v-spacer></v-spacer>
      <v-list-item-icon>
        <v-btn icon v-on:click="openDialog()">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </v-list-item-icon>
    </v-list-item>

    <v-list-group v-for="i in variables" :key="i.Name">
      <template v-slot:activator>
        <v-list-item-icon>
          <v-icon>mdi-variable</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>{{ i.Name }}</v-list-item-title>
        </v-list-item-content>
      </template>

      <v-list-item>
        <v-list-item-icon>
          <v-icon>mdi-tag-multiple</v-icon>
        </v-list-item-icon>
        <v-list-item-content>{{ i.Type }}</v-list-item-content>
        <v-btn icon absolute small right v-on:click="delVariable(i)">
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </v-list-item>
      <v-list-item>
        <v-list-item-icon>
          <v-icon>mdi-view-list</v-icon>
        </v-list-item-icon>
        <v-list-item-content>{{ hexdsp(i.Addr) }}</v-list-item-content>
      </v-list-item>
    </v-list-group>
    <VariableNewDialog @getVariables="getVariables" ref="VariableNewDialog" v-bind:opt="opt" />
  </v-list>
</template>

<script>
import axios from "axios";
import VariableNewDialog from "@/components/VariableNewDialog.vue";
export default {
  props: ["showtext"],
  components: {
    VariableNewDialog,
  },
  data: () => ({
    opt: "",
    variables: [],
  }),
  mounted() {
    this.opt = this.showtext == "观察" ? "read" : "modi";
    this.getVariables();
  },
  methods: {
    openDialog() {
      this.$refs.VariableNewDialog.openDialog();
    },
    hexdsp(i) {
      var h = i.toString(16);
      var l = h.length;
      var z = new Array(9 - l).join("0");
      return "0x" + z + h;
    },
    getVariables() {
      axios.get("/variable-" + this.opt + "/list").then((response) => {
        this.variables = response.data.Variables;
      });
    },
    delVariable(i) {
      axios
        .post("/variable-" + this.opt + "/del", {
          Board: 1,
          Name: i.Name,
          Type: i.Type,
          Addr: i.Addr,
        })
        .then((response) => {
          if (response.data.status == 0) {
            this.getVariables();
            this.$toasted.show("变量删除成功");
          } else if (response.data.status == 22) {
            this.$toasted.error("变量操作时串口错误");
          } else if (response.data.status == 24) {
            this.$toasted.error("删除未添加的变量");
          }
        });
    },
  },
};
</script>
