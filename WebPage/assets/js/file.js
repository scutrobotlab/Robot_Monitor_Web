axios.get('/file/variables')
    .then(function (response) {
        appFileVariables.lists=response.data.Variables
        appFileVariables.searchData=response.data.Variables
    })
    .catch(function (error) {
        console.log(error);
    })

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
                                appFileVariables.searchData=response.data.Variables
                            })
                            .catch(function (error) {
                                console.log(error);
                            })
                    }else if (response.data.status==31){
                        toastShow('未选择文件',1)
                    }else if (response.data.status==32){
                        toastShow('文件写入错误',1)
                    }else if (response.data.status==33){
                        toastShow('文件转换错误',1)
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
        lists:[],
        searchData:[],
        keyword:''
    },
    methods :{
        search: function(event){
            var keyword = this.keyword
            if (keyword) {
                this.searchData = this.lists.filter(function(product) {
                    return Object.keys(product).some(function(key) {
                        return String(product[key]).toLowerCase().indexOf(keyword) > -1
                    })
                })
            }else if(keyword.length==0){
                this.searchData = this.lists
            }else{
                return this.searchData
            }
        },
        variableadd: function(index){
            axios.post('/variable-read/add', {
                    Board: 1,
                    Name: appFileVariables.lists[index].Name,
                    Type: appFileVariables.lists[index].Type,
                    Addr: parseInt(appFileVariables.lists[index].Addr.slice(2),16),
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
                    axios.get('/variable-read/list')
                        .then(function (response) {
                            appVariableList.lists=response.data.Variables
                        })
                        .catch(function (error) {
                            console.log(error);
                        })
                });
        },
        variablemodadd: function(index){
            axios.post('/variable-modi/add', {
                    Board: 1,
                    Name: appFileVariables.lists[index].Name,
                    Type: appFileVariables.lists[index].Type,
                    Addr: parseInt(appFileVariables.lists[index].Addr.slice(2),16),
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
                    axios.get('/variable-modi/list')
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
