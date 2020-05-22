<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <Chart />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Chart from "@/components/Chart.vue";
export default {
  components: {
    Chart
  }
};
</script>
