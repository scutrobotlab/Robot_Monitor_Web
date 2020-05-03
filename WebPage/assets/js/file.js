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
            console.log(param.get('file')); 
            let config = {
                headers:{'Content-Type':'multipart/form-data'}
            };
            axios.post('/file/upload',param,config)
                .then(response=>{
                    console.log(response.data);
                })
                .catch(function (error) {
                    console.log(error);
                })
        }
    }
})
