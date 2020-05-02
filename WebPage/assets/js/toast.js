Vue.use(Toasted)

function toastShow(txt,theme){
    Vue.toasted.show(txt, { 
        theme: theme==0?'toasted-primary':'bubble',
        position: "top-right", 
        duration : 3000
    });
}
