var appSerialList = new Vue({
    el: '#serialport',
    data :{
        selected: '',
        serialLists:[]
    }
})
axios.get('/serial')
    .then(function (response) {
        appSerialList.selected=response.data.Name
    })
    .catch(function (error) {
        console.log(error);
    })
axios.get('/serial/list')
    .then(function (response) {
        appSerialList.serialLists=response.data.Ports
    })
    .catch(function (error) {
        console.log(error);
    })

var appSerialBtn = new Vue({
    el: '#serialbtn',
    methods: {
        openserial: function(event){
            axios.get('/serial/open', {
                    params: {
                        port: appSerialList.selected
                    }
                })
                .then(function (response) {
                    if (response.data.status==0){
                        alert('串口打开成功')
                    }else if (response.data.status==11){
                        alert('无法打开串口')
                    }
                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        closeserial: function(event){
            axios.get('/serial/close')
                .then(function (response) {
                    if (response.data.status==0){
                        alert('串口关闭成功')
                    }else if (response.data.status==12){
                        alert('在未打开串口情况下关闭串口')
                    }
                    else if (response.data.status==13){
                        alert('无法关闭串口')
                    }
                })
                .catch(function (error) {
                    console.log(error);
                })
        }
    }
})
