$(function(){
    $.ajax({
        type: 'GET',
        url: 'list',
        dataType: 'json',
        success : function(data){
            for(var i=0;i<data.Ports.length;i++){
                var result = "<option value='" + data.Ports[i] + "'>" + data.Ports[i] + "</option>";
                $('#serialport').append(result);
            }
        }
    });
    $("#openserial").click(function(event){
        $.ajax({
            type: 'POST',
            data:"port="+$('#serialport').val(), 
            url: 'open',
            success : function(data){
                alert(data);
            }
        });
    });
    $("#closeserial").click(function(event){
        $.ajax({
            type: 'GET',
            url: 'close',
            success : function(data){
                alert(data);
            }
        });
    });
    $("#sendserial").click(function(event){
        $.ajax({
            type: 'POST',
            url: 'act',
            dataType: 'json',
            contentType: 'application/json; charset=utf-8',
            async: false,
            data:JSON.stringify({
                Board:parseInt($('#Board').val()),
                Name:$('#Name').val(),
                Act:parseInt($('#Act').val()),
                Type:$('#Type').val(),
                Addr :parseInt($('#Addr').val(),16),
                Data:parseFloat($('#Data').val()),
            }),
            success : function(data){
                alert(data);
            }
        });
    });
});
