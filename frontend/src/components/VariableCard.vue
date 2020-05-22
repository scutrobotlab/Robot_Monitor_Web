<template>
  <v-card>
    <v-toolbar dense>
      <v-toolbar-title>变量·{{showtext}}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon v-on:click="openDialog()">
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </v-toolbar>
    <v-card-text>
      <v-expansion-panels focusable>
        <v-expansion-panel
          v-for="i in variables"
          :key="i.Name"
        >
          <v-expansion-panel-header>
            {{i.Name}}
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-icon>mdi-tag-multiple</v-icon>{{i.Type}}
            <v-icon>mdi-view-list</v-icon>{{hexdsp(i.Addr)}}
            <v-btn icon absolute small right v-on:click="delVariable(i)"><v-icon>mdi-delete</v-icon></v-btn>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card-text>
    <VariableNewDialog @getVariables="getVariables" ref="VariableNewDialog" v-bind:opt="opt"/>
    <Notice ref="notice"/>
  </v-card>
</template>

<script>
  import axios from "axios"
  import VariableNewDialog from "@/components/VariableNewDialog.vue"
  import Notice from "@/components/Notice.vue"
  export default {
    props: [
      "showtext"
    ],
    components: {
      VariableNewDialog,
      Notice
    },
    data: () => ({
      opt: '',
      variables: []
    }),
    mounted() {
      this.opt = this.showtext=='观察'?'read':'modi'
      this.getVariables()
    },
    methods: {
      openDialog(){
        this.$refs.VariableNewDialog.openDialog()
      },
      hexdsp(i){
        var h = i.toString(16)
        var l = h.length
        var z = new Array(9-l).join("0")
        return "0x"+z+h
      },
      getVariables(){
        axios.get('/variable-'+this.opt+'/list')
        .then(response =>{
          this.variables = response.data.Variables
        })
      },
      delVariable(i){
        axios.post('/variable-'+this.opt+'/del', {
            Board: 1,
            Name: i.Name,
            Type: i.Type,
            Addr: i.Addr,
        })
        .then(response =>{
          if (response.data.status==0){
            this.getVariables()
            this.$refs.notice.show('变量删除成功',0)
          }else if (response.data.status==22){
            this.$refs.notice.show('变量操作时串口错误',1)
          }
          else if (response.data.status==24){
            this.$refs.notice.show('删除未添加的变量',1)
          }
        })
      }
    }
  }
</script>
