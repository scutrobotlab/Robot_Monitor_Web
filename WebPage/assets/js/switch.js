$.fn.bootstrapSwitch.defaults.onText = '是';
$.fn.bootstrapSwitch.defaults.offText = '否';
$.fn.bootstrapSwitch.defaults.onColor = 'info';
$.fn.bootstrapSwitch.defaults.offColor = 'danger';
$.fn.bootstrapSwitch.defaults.size = 'mini';

axios.get('/file/config')
    .then(function (response) {
        $("[name='checkbox-sda']").bootstrapSwitch('state',response.data.IsSaveDataAddr,true)
        $("[name='checkbox-svm']").bootstrapSwitch('state',response.data.IsSaveVariablesToMod,true)
        $("[name='checkbox-svr']").bootstrapSwitch('state',response.data.IsSaveVariablesToRead,true)
    })
    .catch(function (error) {
        console.log(error);
    })

$("[name='checkbox-sda']").bootstrapSwitch({
    onSwitchChange:function(event,state){
        axios.get('/file/config', {
            params: {
                sda: state
            }
        })
    }
});
$("[name='checkbox-svr']").bootstrapSwitch({
    onSwitchChange:function(event,state){
        axios.get('/file/config', {
            params: {
                svr: state
            }
        })
    }
});
$("[name='checkbox-svm']").bootstrapSwitch({
    onSwitchChange:function(event,state){
        axios.get('/file/config', {
            params: {
                svm: state
            }
        })
    }
});
