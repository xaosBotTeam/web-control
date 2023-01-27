var sign_in = new Vue({
    el: '#form-login',
    data() {
        return {
            login:'',
            password: ''
        }
    },
    methods: {
        formSubmit(e) {
            console.log(this)
            e.preventDefault();
            axios.post(`/auth`, {
                login: this.login,
                password: this.password,
            })
                .then(function () {
                    window.location.href = "control-panel.html"
                })
        }
    },
})