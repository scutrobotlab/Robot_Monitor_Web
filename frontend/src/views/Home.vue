<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <ChartCard />
      </v-col>
      <PanelCard />
    </v-row>
  </v-container>
</template>

<script>
import ChartCard from "@/components/ChartCard.vue";
import PanelCard from "@/components/PanelCard.vue";
export default {
  components: {
    ChartCard,
    PanelCard
  }
};
</script>
