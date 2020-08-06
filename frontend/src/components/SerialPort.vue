<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>串口</v-list-item-title>
    </v-list-item>
    <v-list-item>
      <v-list-item-action>
        <v-switch v-model="status" v-on:change="optSerial()" inset></v-switch>
      </v-list-item-action>
      <v-select
        :items="serialList"
        v-model="serial"
        v-on:click="getSerialList()"
        v-bind:disabled="status"
        label="选择串口"
      ></v-select>
    </v-list-item>
  </v-list>
</template>

<script>
import axios from "axios";
export default {
  data: () => ({
    status: false,
    serial: null,
    serialList: [],
  }),
  mounted() {
    this.getSerialList();
    this.getSerial();
  },
  methods: {
    getSerialList() {
      axios.get("/serial/list").then((response) => {
        this.serialList = response.data.Ports;
      });
    },
    getSerial() {
      axios.get("/serial").then((response) => {
        this.serial = response.data.Name;
        if (this.serial) {
          this.status = true;
        }
      });
    },
    optSerial() {
      if (this.status) {
        axios.get("/serial/open?port=" + this.serial).then((response) => {
          if (response.data.status == 0) {
            this.$toasted.show("串口打开成功");
          } else if (response.data.status == 1) {
            this.$toasted.error("未选择串口");
            this.status = false;
          } else if (response.data.status == 11) {
            this.$toasted.error("无法打开串口");
            this.status = false;
          }
        });
      } else {
        axios.get("/serial/close").then((response) => {
          if (response.data.status == 0) {
            this.$toasted.show("串口关闭成功");
          } else if (response.data.status == 12) {
            this.$toasted.error("在未打开串口情况下关闭串口");
            this.status = true;
          } else if (response.data.status == 13) {
            this.$toasted.error("无法关闭串口");
            this.status = true;
          }
        });
      }
    },
  },
};
</script>
