var ws = new WebSocket("ws://"+window.location.host+"/ws");
ws.onopen = function(evt) {
    toastShow("连接成功",0);
}
ws.onclose = function(evt) {
    toastShow("连接断开",1);
    ws = null;
}
ws.onmessage = function(evt) {
    praseWS(evt.data);
}
ws.onerror = function(evt) {
    console.log("ERROR: " + evt.data);
}

$("[name='checkbox-ws']").bootstrapSwitch({
    onText: '启动',
    offText: '停止'
})

function praseWS(data){
    if(data!=""){
        const jsonWS = JSON.parse(data);
        for(i in jsonWS.DataPack){
            chartData[i].push({x:jsonWS.DataPack[i].Tick,y:jsonWS.DataPack[i].Data});
        }
        if($("[name='checkbox-ws']").prop("checked")){
            chart.update();
        }
    }
}
