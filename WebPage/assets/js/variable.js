function hexdsp(i){
    var h = i.toString(16)
    var l = h.length
    var z = new Array(9-l).join("0")
    return "0x"+z+h
}

var appVariableList = new Vue({
    el: '#variablelist',
    data :{
        lists:[]
    }
})
axios.get('/variable')
    .then(function (response) {
        appVariableList.lists=response.data.Variables
    })
    .catch(function (error) {
        console.log(error);
    })

var appVariableAdd = new Vue({
    el: '#variableadd',
    data :{
        Board: 1,
        Name: '',
        selected: '',
        Types:[],
        Addr: 0,
    },
    methods :{
        variableadd: function(event){
            axios.post('/variable/add', {
                    Board: 1,
                    Name: appVariableAdd.Name,
                    Type: appVariableAdd.selected,
                    Addr: parseInt(appVariableAdd.Addr,16),
                })
                .then(function (response) {
                    if (response.data.status==0){
                        alert('变量添加成功')
                    }else if (response.data.status==22){
                        alert('变量操作时串口错误')
                    }
                    else if (response.data.status==23){
                        alert('重复添加变量')
                    }
                })
                .catch(function (error) {
                    console.log(error);
                })
                .then(function () {
                    axios.get('/variable')
                        .then(function (response) {
                            appVariableList.lists=response.data.Variables
                        })
                        .catch(function (error) {
                            console.log(error);
                        })
                });
        }
    }
})
axios.get('/variable/types')
    .then(function (response) {
        appVariableAdd.Types=response.data.Types
    })
    .catch(function (error) {
        console.log(error);
    })
