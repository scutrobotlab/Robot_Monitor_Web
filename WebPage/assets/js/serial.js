var appSerial = new Vue({
    el: '#serial',
    data :{
        selected: '',
        serialLists:[]
    },
    methods: {
        refreshserial: function(event){
            axios.get('/serial/list')
                .then(function (response) {
                    appSerial.serialLists=response.data.Ports
                })
                .catch(function (error) {
                    console.log(error);
                })
        }
    }
})

axios.get('/serial')
    .then(function (response) {
        appSerial.selected=response.data.Name
        $("[name='checkbox-serial']").bootstrapSwitch('state',response.data.Name!="",true)
    })
    .catch(function (error) {
        console.log(error);
    })
axios.get('/serial/list')
    .then(function (response) {
        appSerial.serialLists=response.data.Ports
    })
    .catch(function (error) {
        console.log(error);
    })


$("[name='checkbox-serial']").bootstrapSwitch({
    onText: '开',
    offText: '关',
    onSwitchChange:function(event,state){
        if(state){
            axios.get('/serial/open', {
                params: {
                    port: $("#serialport").val()
                }
            })
                .then(function (response) {
                    if (response.data.status==0){
                        toastShow('串口打开成功',0)
                    }else if (response.data.status==1){
                        toastShow('未选择串口',1)
                        $("[name='checkbox-serial']").bootstrapSwitch('state',false,true)
                    }else if (response.data.status==11){
                        toastShow('无法打开串口',1)
                        $("[name='checkbox-serial']").bootstrapSwitch('state',false,true)
                    }
                })
                .catch(function (error) {
                    console.log(error);
                })
        }else{
            axios.get('/serial/close')
                .then(function (response) {
                    if (response.data.status==0){
                        toastShow('串口关闭成功',0)
                    }else if (response.data.status==12){
                        toastShow('在未打开串口情况下关闭串口',1)
                        $("[name='checkbox-serial']").bootstrapSwitch('state',true,true)
                    }
                    else if (response.data.status==13){
                        toastShow('无法关闭串口',1)
                        $("[name='checkbox-serial']").bootstrapSwitch('state',true,true)
                    }
                })
                .catch(function (error) {
                    console.log(error);
                })
        }
    }
});
