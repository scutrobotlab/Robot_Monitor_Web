const el = document.getElementById('chart');
const chartData = [];
const chart = new TimeChart(el, {
    series: [{ data:chartData }],
});
