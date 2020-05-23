<template>
  <v-card>
    <div id="chart" style="width: 100%; height: 420px;"></div>
  </v-card>
</template>
<script>
import timechart from "timechart";
import colors from "vuetify/lib/util/colors";
export default {
  data: () => ({
    ws: null,
    chart: null,
    indexColor: 0,
    lineColors: {
      light: [
        colors.red.lighten1,
        colors.green.lighten1,
        colors.orange.lighten1,
        colors.purple.lighten1,
        colors.indigo.lighten1,
        colors.teal.lighten1,
        colors.pink.lighten1
      ],
      dark: [
        colors.red.darken4,
        colors.green.darken4,
        colors.orange.darken4,
        colors.purple.darken4,
        colors.indigo.darken4,
        colors.teal.darken4,
        colors.pink.darken4
      ]
    }
  }),
  computed: {
    isDark() {
      return this.$vuetify.theme.dark;
    }
  },
  created() {
    this.initWS();
  },
  destroyed() {
    this.ws.close();
  },
  mounted() {
    const el = document.getElementById("chart");
    this.chart = new timechart(el, {
      baseTime: Date.now() - performance.now(),
      series: [],
      xRange: { min: 0, max: 20 * 1000 },
      realTime: true
    });
  },
  watch: {
    isDark: function() {
      if (this.isDark) {
        for (var i in this.chart.options.series) {
           this.chart.options.series[i].color = this.lineColors.dark[i];
        }
      } else {
        for (i in this.chart.options.series) {
           this.chart.options.series[i].color = this.lineColors.light[i];
        }
      }
      this.chart.update();
    }
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
      if (data != "") {
        const jsonWS = JSON.parse(data);
        for (var i in jsonWS.DataPack) {
          const fi = this.chart.options.series.findIndex(
            a => a.name == jsonWS.DataPack[i].Name
          );
          if (fi > -1) {
            this.chart.options.series[fi].data.push({
              x: jsonWS.DataPack[i].Tick,
              y: jsonWS.DataPack[i].Data
            });
          } else {
            const color = this.selectColor();
            this.chart.options.series.push({
              name: jsonWS.DataPack[i].Name,
              color: color,
              data: [{ x: jsonWS.DataPack[i].Tick, y: jsonWS.DataPack[i].Data }]
            });
            this.iColor++;
          }
        }
        this.chart.update();
      }
    }
  }
};
</script>
