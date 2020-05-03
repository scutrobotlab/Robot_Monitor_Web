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
