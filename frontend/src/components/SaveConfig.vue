<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>保存选项</v-list-item-title>
    </v-list-item>
    <v-list-item v-for="i in config" :key="i.n">
      <v-switch v-model="i.v" :label="i.t" v-on:change="updateConfig(i.n,i.v)" inset></v-switch>
    </v-list-item>
  </v-list>
</template>

<script>
import axios from "axios";
export default {
  data: () => ({
    config: [
      {
        t: "变量列表",
        n: "sda",
        v: true
      },
      {
        t: "观察变量",
        n: "svr",
        v: true
      },
      {
        t: "修改变量",
        n: "svm",
        v: true
      }
    ]
  }),
  mounted() {
    this.getConfig();
  },
  methods: {
    getConfig() {
      axios.get("/file/config").then(response => {
        this.config[0].v = response.data.IsSaveDataAddr;
        this.config[1].v = response.data.IsSaveVariableRead;
        this.config[2].v = response.data.IsSaveVariableModi;
      });
    },
    updateConfig(n, v) {
      axios.get("/file/config?" + n + "=" + v);
    }
  }
};
</script>
