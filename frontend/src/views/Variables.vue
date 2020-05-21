<template>
  <v-container>
    <v-row>

      <v-col cols="12" sm="6" md="6">
        <VariableCard showtext='观察'/>
      </v-col>

      <v-col cols="12" sm="6" md="6">
        <VariableCard showtext='修改'/>
      </v-col>

      <v-btn color="secondary" dark absolute bottom right fab v-on:click="openDialog()"><v-icon>mdi-plus</v-icon></v-btn>
      <VariableAllDialog ref="VariableAllDialog"/>
    </v-row>
  </v-container>
</template>

<script>
  import VariableAllDialog from "@/components/VariableAllDialog.vue"
  import VariableCard from "@/components/VariableCard.vue"
  export default {
    components:{
      VariableCard,
      VariableAllDialog
    },
    methods:{
      openDialog(){
        this.$refs.VariableAllDialog.openDialog()
      }
    }
  }
</script>
