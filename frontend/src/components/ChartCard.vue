<template>
  <v-card>
    <div id="chart" style="width: 100%; height: 420px;"></div>
  </v-card>
</template>
<script>
import timechart from "timechart";
export default {
  data: () => ({
    ws: null,
    chart: null,
    iColor: 0,
    colors: [
      "#F44336",
      "#9C27B0",
      "#3F51B5",
      "#00BCD4",
      "#4CAF50",
      "#FF9800",
      "#795548"
    ]
  }),
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
  methods: {
    initWS() {
      this.ws = new WebSocket("ws://" + window.location.host + "/ws");
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
            this.chart.options.series.push({
              name: jsonWS.DataPack[i].Name,
              color: this.colors[this.iColor],
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
