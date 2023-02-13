var sign_in = new Vue({
    el: '#change-password',
    data() {
        return {
            new_password1: '',
            new_password2: '',
            err_msg:''
        }
    },
    methods: {
        formSubmit(e) {
            e.preventDefault();

            if(this.new_password1!==this.new_password2){
                this. err_msg = "Пароли не совпадают"
                return
            }
            else{
                this. err_msg = ""
            }

            axios.post(`/resetPassword`, {
                password: this.new_password1,
            }).catch(error => {if(error.response.status==403){
                    window.location.href = "sign-in.js.html"
                }
                    console.log(error)});;
        },
    },
})