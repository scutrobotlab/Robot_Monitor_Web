<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <ChartCard />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import ChartCard from "@/components/ChartCard.vue";
export default {
  components: {
    ChartCard
  }
};
</script>
