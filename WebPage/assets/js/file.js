var appFileUpload = new Vue({
    el: '#fileupload',
    data :{
        file:{name:'变量表文件'}
    },
    methods:{
        choose: function(event){
            appFileUpload.file=event.target.files[0]
        },
        upload: function(event){
            let param = new FormData();
            param.append('file',appFileUpload.file);
            let config = {
                headers:{'Content-Type':'multipart/form-data'}
            };
            axios.post('/file/upload',param,config)
                .then(response=>{
                    if (response.data.status==0){
                        toastShow('文件上传成功',0)
                        axios.get('/file/variables')
                            .then(function (response) {
                                appFileVariables.lists=response.data.Variables
                            })
                            .catch(function (error) {
                                console.log(error);
                            })
                    }else if (response.data.status==31){
                        toastShow('未选择文件',1)
                    }else if (response.data.status==32){
                        toastShow('文件写入错误',1)
                    }
                })
                .catch(function (error) {
                    console.log(error);
                })
        }
    }
})

var appFileVariables = new Vue({
    el: '#filevariables',
    data :{
        lists:[]
    },
    methods :{
        variableadd: function(event){
            axios.post('/variable/add', {
                    Board: 1,
                    Name: appFileVariables.Name,
                    Type: appFileVariables.Type,
                    Addr: parseInt(appFileVariables.Addr.slice(2),16),
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
        },
        variablemodadd: function(event){
            axios.post('/variable/modadd', {
                    Board: 1,
                    Name: appFileVariables.Name,
                    Type: appFileVariables.Type,
                    Addr: parseInt(appFileVariables.Addr.slice(2),16),
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
                    axios.get('/variable/modlist')
                        .then(function (response) {
                            appVariableModList.lists=response.data.Variables
                        })
                        .catch(function (error) {
                            console.log(error);
                        })
                });
            }
    }
})
