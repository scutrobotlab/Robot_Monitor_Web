$(function(){
    $.ajax({
        type: 'GET',
        url: '/serial/list',
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
            url: '/serial/open',
            success : function(data){
                alert(data);
            }
        });
    });
    $("#closeserial").click(function(event){
        $.ajax({
            type: 'GET',
            url: '/serial/close',
            success : function(data){
                alert(data);
            }
        });
    });
    $("#sendserial").click(function(event){
        $.ajax({
            type: 'POST',
            url: '/variable/add',
            dataType: 'json',
            contentType: 'application/json; charset=utf-8',
            async: false,
            data:JSON.stringify({
                Board:parseInt($('#variable-board').val()),
                Name:$('#variable-name').val(),
                Type:$('#variable-type').val(),
                Addr :parseInt($('#variable-addr').val(),16),
                Data:parseFloat($('#variable-data').val()),
            }),
            success : function(data){
                alert(data);
            }
        });
        $("#NewVariable").modal('hide');
    });
});
