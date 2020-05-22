<template>
  <v-container>
    <v-row>
      <v-col cols="12" sm="6" md="6">
        <SerialPort />
      </v-col>

      <v-col cols="12" sm="6" md="6">
        <SaveConfig />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import SerialPort from "@/components/SerialPort.vue";
import SaveConfig from "@/components/SaveConfig.vue";
export default {
  components: {
    SaveConfig,
    SerialPort
  }
};
</script>
