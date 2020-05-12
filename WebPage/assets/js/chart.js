const el = document.getElementById('chart');
const chartData = [[],[]];
const chart = new TimeChart(el, {
    series: [
        {
            name:'a',
            data:chartData[0]
        },
        {
            name:'b',
            color: 'red',
            data:chartData[1]
        }
    ]
});
