window.addEventListener("load", function(evt) {
    var ws;
    $("#openWS").click (function(event){
        if (ws) {
            return false;
        }
        ws = new WebSocket("ws://"+document.domain+":8080/ws");
        ws.onopen = function(evt) {
            toastShow("连接成功",0);
        }
        ws.onclose = function(evt) {
            ws = null;
        }
        ws.onmessage = function(evt) {
            var jsonWS = JSON.parse(evt.data);
            var l = chartData.length;
            var d = {'x':l,'y':jsonWS.DataPack[0].Data};
            chartData.push(d);
            chart.update();
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