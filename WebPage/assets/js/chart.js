const el = document.getElementById('chart');
const chart = new TimeChart(el, {
    baseTime: Date.now() - performance.now(),
    series: [],
    xRange: { min: 0, max: 20 * 1000 },
    realTime: true
});
const colors=new Array('#F44336','#9C27B0','#3F51B5','#00BCD4','#4CAF50','#FF9800','#795548');
