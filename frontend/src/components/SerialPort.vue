<template>
  <v-card>
    <v-card-title>串口</v-card-title>
    <v-card-text>
      <v-row>
        <v-col cols="9">
          <v-select
            prepend-icon="mdi-serial-port"
            :items="serialList"
            v-model="serial"
            v-on:click="getSerialList()"
            v-bind:disabled="status"
            label="选择串口"
            solo
          >
          </v-select>
        </v-col>
        <v-col cols="3">
          <v-switch
            v-model="status"
            v-on:change="optSerial()"
            inset
          >
          </v-switch>
        </v-col>
        <Notice ref="notice"/>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script>
  import axios from "axios";
  import Notice from "@/components/Notice.vue"
  export default {
    components:{
      Notice
    },
    data: () => ({
      status:false,
      serial:null,
      serialList:[]
    }),
    mounted() {
      this.getSerialList()
      this.getSerial()
    },
    methods :{
      getSerialList(){
        axios.get('/serial/list')
        .then(response =>{
          this.serialList = response.data.Ports
        })
      },
      getSerial(){
        axios.get('/serial')
        .then(response =>{
            this.serial= response.data.Name
            if (this.serial){
              this.status=true
            }
        })
      },
      optSerial(){
        if(this.status){
          axios.get('/serial/open?port=' + this.serial)
          .then(response =>{
            if(response.data.status==0){
              this.$refs.notice.show('串口打开成功',0)
            }else if(response.data.status==1){
              this.$refs.notice.show('未选择串口',1)
              this.status=false
            }else if(response.data.status==11){
              this.$refs.notice.show('无法打开串口',1)
              this.status=false
            }
          })
        }else{
          axios.get('/serial/close')
          .then(response =>{
            if (response.data.status==0){
              this.$refs.notice.show('串口关闭成功',0)
            }else if (response.data.status==12){
              this.$refs.notice.show('在未打开串口情况下关闭串口',1)
              this.status=true
            }else if (response.data.status==13){
              this.$refs.notice.show('无法关闭串口',1)
              this.status=true
            }
          })
        }
      }
    }
  }
</script>
