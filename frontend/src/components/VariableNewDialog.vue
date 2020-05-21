<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">添加变量</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field label="板子代号" type="number" hint="保持默认为1即可" required v-model="Board"></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field label="变量名" type="text" required v-model="Name"></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-select
                  :items="types"
                  label="变量类型"
                  required
                  v-model="Type"
                ></v-select>
              </v-col>
              <v-col cols="12">
                <v-text-field label="变量地址" type="text" hint="形如2000ab78" required v-model="Addr"></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="dialog = false">关闭</v-btn>
          <v-btn color="blue darken-1" text @click="addVariable()">添加</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <Notice ref="notice"/>
  </v-row>
</template>

<script>
  import axios from "axios";
  import Notice from "@/components/Notice.vue"
  export default {
    components:{
      Notice
    },
    props: [
      'opt'
    ],
    data: () => ({
      dialog: false,
      types: [],
      Board: 1,
      Name: '',
      Type: '',
      Addr: '',
    }),
    mounted(){
      axios.get('/variable/types')
      .then(response =>{
        this.types = response.data.Types
      })
    },
    methods:{
      openDialog(){
        this.dialog = true
      },
      addVariable(){
        axios.post('/variable-'+this.opt+'/add', {
            Board: 1,
            Name: this.Name,
            Type: this.Type,
            Addr: parseInt(this.Addr,16),
        })
        .then(response =>{
          if (response.data.status==0){
            this.dialog = false
            this.$refs.notice.show('变量添加成功',0)
          }else if (response.data.status==22){
            this.$refs.notice.show('变量操作时串口错误',1)
          }else if (response.data.status==23){
            this.$refs.notice.show('重复添加变量',1)
          }
        })
      },
    }
  }
</script>
