<template>
  <v-container>
    <ChartCard />
    <PanelCard />
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
