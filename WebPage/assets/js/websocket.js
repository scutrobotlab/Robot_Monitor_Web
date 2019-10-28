window.addEventListener("load", function(evt) {
    var ws;
    $("#openWS").click (function(event){
        if (ws) {
            return false;
        }
        ws = new WebSocket("ws://"+document.domain+":8080/ws");
        ws.onopen = function(evt) {
            alert("Websocket connection established");
        }
        ws.onclose = function(evt) {
            ws = null;
        }
        ws.onmessage = function(evt) {
            var jsonWS = JSON.parse(evt.data);
            console.log(jsonWS.DataPack[0].Data)
            var l =  myChart.data.datasets[0].data.length;
            var d = {'x':l,'y':jsonWS.DataPack[0].Data}
            myChart.data.labels.push(l)
            myChart.data.datasets[0].data.push(d)
            myChart.update()
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    });
    $("#closeWS").click (function(event) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    });
});