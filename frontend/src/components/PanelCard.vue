<template>
  <div class="text-center">
    <v-menu
      v-model="menu"
      :close-on-content-click="false"
      transition="slide-x-transition"
      top
      offset-y
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          class="ma-8"
          color="secondary"
          dark
          absolute
          bottom
          left
          fab
          v-on="on"
          v-bind="attrs"
        >
          <v-icon>mdi-iframe-variable</v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>调参面板</v-card-title>
        <v-list dense>
          <v-list-item-group color="primary">
            <v-list-item v-for="i in variables" :key="i.Name">
              <v-list-item-content>
                <v-text-field
                  dense
                  v-model="i.Data"
                  v-bind:label="i.Name"
                  v-on:keyup.enter="modiVariable(i)"
                ></v-text-field>
              </v-list-item-content>
              <v-list-item-icon>
                <v-btn icon v-on:click="modiVariable(i)">
                  <v-icon>mdi-send</v-icon>
                </v-btn>
              </v-list-item-icon>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-card>
    </v-menu>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data: () => ({
    menu: false,
    variables: []
  }),
  mounted() {
    this.getVariables();
  },
  methods: {
    openMenu() {
      this.menu = true;
    },
    getVariables() {
      axios.get("/variable-modi/list").then(response => {
        this.variables = response.data.Variables;
      });
    },
    modiVariable(i) {
      axios
        .post("/variable-modi/mod", {
          Board: 1,
          Name: i.Name,
          Type: i.Type,
          Addr: i.Addr,
          Data: parseFloat(i.Data)
        })
        .then(response => {
          if (response.data.status == 0) {
            this.$toasted.show("变量修改成功");
          } else if (response.data.status == 22) {
            this.$toasted.error("变量操作时串口错误");
          }
        });
    }
  }
};
</script>
