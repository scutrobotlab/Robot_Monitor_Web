<template>
  <div>
    <div id="chart" style="width: 100%; height: 600px;"></div>
    <Notice ref="notice" />
  </div>
</template>

<script>
import timechart from "timechart";
import Notice from "@/components/Notice.vue";
export default {
  components: {
    Notice
  },
  data: () => ({
    el: null,
    chart: null,
    cColors: 0,
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
  mounted() {
    this.el = document.getElementById("chart");
    this.chart = new timechart(this.el, {
      baseTime: Date.now() - performance.now(),
      series: [],
      xRange: { min: 0, max: 20 * 1000 },
      realTime: true
    });
    var ws = new WebSocket("ws://" + window.location.host + "/ws");
    ws.onopen = function() {
      this.$refs.notice.show("连接成功", 0);
    };
    ws.onclose = function() {
      this.$refs.notice.show("连接断开", 1);
      ws = null;
    };
    ws.onmessage = function(evt) {
      this.praseWS(evt.data);
    };
    ws.onerror = function(evt) {
      console.log("ERROR: " + evt.data);
    };
  },
  methods: {
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
              color: this.colors[this.cColors],
              data: [{ x: jsonWS.DataPack[i].Tick, y: jsonWS.DataPack[i].Data }]
            });
            this.cColors++;
          }
        }
        this.chart.update();
      }
    }
  }
};
</script>
