var sign_in = new Vue({
    el: '#form-login',
    data() {
        return {
            login:'',
            password: '',
            err_msg:''
        }
    },
    methods: {
        formSubmit(e) {
            e.preventDefault();
            axios.post(`/auth`, {
                login: this.login,
                password: this.password,
            })
                .then(function () {
                    window.location.href = "control-panel.html"
                })
                .catch(this.err_msg='Неверный логин или пароль')
        }
    },
})