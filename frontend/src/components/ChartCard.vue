<template>
  <v-card style="overflow-y: hidden;">
    <div ref="chart" style="width: 100%; height: 77vh;"></div>
  </v-card>
</template>
<script>
import timechart from "timechart";
import colors from "vuetify/lib/util/colors";
export default {
  data: () => ({
    ws: null,
    chart: null,
    indexColor: -1,
    lineColors: {
      light: [
        colors.red.lighten1,
        colors.green.lighten1,
        colors.orange.lighten1,
        colors.purple.lighten1,
        colors.indigo.lighten1,
        colors.teal.lighten1,
        colors.pink.lighten1,
      ],
      dark: [
        colors.red.darken4,
        colors.green.darken4,
        colors.orange.darken4,
        colors.purple.darken4,
        colors.indigo.darken4,
        colors.teal.darken4,
        colors.pink.darken4,
      ],
    },
  }),
  computed: {
    isDark() {
      return this.$vuetify.theme.dark;
    },
  },
  created() {
    this.initWS();
  },
  destroyed() {
    this.ws.close();
  },
  mounted() {
    this.chart = new timechart(this.$refs.chart, {
      baseTime: Date.now(),
      series: [],
      xRange: { min: 0, max: 20 * 1000 },
      realTime: true,
      zoom: {
        x: {
          autoRange: true,
          minDomainExtent: 50,
        },
        y: {
          autoRange: true,
        },
      },
    });
  },
  watch: {
    isDark: function () {
      if (this.isDark) {
        for (var i in this.chart.options.series) {
          this.chart.options.series[i].color = this.lineColors.dark[i];
        }
        this.$refs.chart.querySelector("chart-legend").style.backgroundColor =
          "black";
      } else {
        for (i in this.chart.options.series) {
          this.chart.options.series[i].color = this.lineColors.light[i];
        }
        this.$refs.chart.querySelector("chart-legend").style.backgroundColor =
          "white";
      }
      this.chart.update();
    },
  },
  methods: {
    initWS() {
      this.ws = new WebSocket(
        (document.location.protocol == "https:" ? "wss" : "ws") +
          "://" +
          window.location.host +
          "/ws"
      );
      this.ws.onopen = this.WSonopen;
      this.ws.onclose = this.WSclose;
      this.ws.onmessage = this.WSonmessage;
      this.ws.onerror = this.WSonerror;
    },
    WSonopen() {
      this.$toasted.show("连接成功");
    },
    WSclose() {
      this.$toasted.error("连接断开");
    },
    WSonmessage(evt) {
      this.praseWS(evt.data);
    },
    WSonerror(evt) {
      console.log("ERROR: " + evt.data);
    },
    selectColor() {
      this.indexColor++;
      this.indexColor == this.lineColors.light.length ? 0 : this.indexColor;
      if (this.$vuetify.theme.dark) {
        return this.lineColors.dark[this.indexColor];
      } else {
        return this.lineColors.light[this.indexColor];
      }
    },
    praseWS(data) {
      if (!data) {
        return;
      }

      const jsonWS = JSON.parse(data);
      for (const dp of jsonWS.DataPack) {
        let series = this.chart.options.series.find(
          (a) => a.name == dp.Name
        );
        if (!series) {
          const color = this.selectColor();
          series = {
            name: dp.Name,
            color: color,
            data: [],
          };
          this.chart.options.series.push(series);
          this.iColor++;
        }
        series.data.push({
          x: dp.Tick / 1000,
          y: dp.Data,
        });
      }
      this.chart.update();
    },
  },
};
</script>
