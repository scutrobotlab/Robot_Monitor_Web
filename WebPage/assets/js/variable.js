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
    },
    methods :{
        variabledel: function(index){
            axios.post('/variable/del', {
                Board: 1,
                Name: appVariableList.lists[index].Name,
                Type: appVariableList.lists[index].Type,
                Addr: appVariableList.lists[index].Addr,
            })
            .then(function (response) {
                if (response.data.status==0){
                    toastShow('变量删除成功',0)
                }else if (response.data.status==22){
                    toastShow('变量操作时串口错误',1)
                }
                else if (response.data.status==24){
                    toastShow('删除未添加的变量',1)
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
                        toastShow('变量添加成功',0)
                    }else if (response.data.status==22){
                        toastShow('变量操作时串口错误',1)
                    }
                    else if (response.data.status==23){
                        toastShow('重复添加变量',1)
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
