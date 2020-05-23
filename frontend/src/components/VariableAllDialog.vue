<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
      <v-card>
        <v-toolbar dark color="primary">
          <v-toolbar-items>
            <v-btn icon dark v-on:click="closeDialog">
              <v-icon>mdi-close</v-icon>
            </v-btn>
          </v-toolbar-items>
          <v-toolbar-title>变量列表</v-toolbar-title>
        </v-toolbar>
        <v-row>
          <v-col cols="4">
            <v-file-input label="上传变量地址表文件" v-model="file"></v-file-input>
          </v-col>
          <v-col cols="8">
            <v-text-field clearable placeholder="搜索变量" prepend-icon="mdi-magnify" v-model="keyword"></v-text-field>
          </v-col>
        </v-row>

        <v-simple-table dense fixed-header>
          <template v-slot:default>
            <thead>
              <tr>
                <th>名称</th>
                <th>类型</th>
                <th>地址</th>
                <th>观察</th>
                <th>修改</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="list in searchData" :key="list.Name">
                <td>{{list.Name}}</td>
                <td>{{list.Type}}</td>
                <td>{{list.Addr}}</td>
                <td>
                  <v-btn icon v-on:click="variableReadAdd(list)">
                    <v-icon>mdi-plus</v-icon>
                  </v-btn>
                </td>
                <td>
                  <v-btn icon v-on:click="variableModiAdd(list)">
                    <v-icon>mdi-plus</v-icon>
                  </v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import axios from "axios";
export default {
  data: () => ({
    dialog: false,
    file: null,
    lists: [],
    searchData: [],
    keyword: ""
  }),
  watch: {
    file: function() {
      let file = this.file;
      let param = new FormData();
      param.append("file", file);
      let config = {
        headers: { "Content-Type": "multipart/form-data" }
      };
      axios.post("/file/upload", param, config).then(response => {
        if (response.data.status == 0) {
          this.getVariableList();
          this.$toasted.show("文件上传成功");
        } else if (response.data.status == 31) {
          this.$toasted.error("未选择文件");
        } else if (response.data.status == 32) {
          this.$toasted.error("文件写入错误");
        } else if (response.data.status == 33) {
          this.$toasted.error("文件转换错误");
        }
      });
    },
    keyword: function() {
      var keyword = this.keyword;
      if (keyword) {
        this.searchData = this.lists.filter(function(product) {
          return Object.keys(product).some(function(key) {
            return (
              String(product[key])
                .toLowerCase()
                .indexOf(keyword) > -1
            );
          });
        });
      } else if (keyword.length == 0) {
        this.searchData = this.lists;
      } else {
        return this.searchData;
      }
    }
  },
  mounted() {
    this.getVariableList();
  },
  methods: {
    openDialog() {
      this.dialog = true;
    },
    closeDialog() {
      this.$emit("getAllV");
      this.dialog = false;
    },
    getVariableList() {
      axios.get("/file/variables").then(response => {
        this.lists = response.data.Variables;
        this.searchData = response.data.Variables;
      });
    },
    variableReadAdd(i) {
      axios
        .post("/variable-read/add", {
          Board: 1,
          Name: i.Name,
          Type: i.Type,
          Addr: parseInt(i.Addr, 16)
        })
        .then(response => {
          if (response.data.status == 0) {
            this.$toasted.show("变量添加成功");
          } else if (response.data.status == 22) {
            this.$toasted.error("变量操作时串口错误");
          } else if (response.data.status == 23) {
            this.$toasted.error("重复添加变量");
          }
        });
    },
    variableModiAdd(i) {
      axios
        .post("/variable-modi/add", {
          Board: 1,
          Name: i.Name,
          Type: i.Type,
          Addr: parseInt(i.Addr, 16)
        })
        .then(response => {
          if (response.data.status == 0) {
            this.$toasted.show("变量添加成功");
          } else if (response.data.status == 22) {
            this.$toasted.error("变量操作时串口错误");
          } else if (response.data.status == 23) {
            this.$toasted.error("重复添加变量");
          }
        });
    }
  }
};
</script>
