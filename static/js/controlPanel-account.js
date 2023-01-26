Vue.component('account-item', {
    methods: {
        formSubmit(e) {
            console.log(this)
            e.preventDefault();
            axios.post(`/api/account/${this.account.ID}`, {
                Sliv: this.account.Sliv,
            })
        }
    },

    props: ['account'],

    template:
        '<div class="col">' +
        '   <form @submit="formSubmit">'+
        '       <h1>{{account.FriendlyName}}</h1>' +
        '            <div class="form-check form-switch">' +
        '                <input class="form-check-input" type="checkbox" role="switch" v-model="account.Sliv">' +
        '                <label class="form-check-label">Слив</label>' +
        '            </div>' +
        '       <button class="btn btn-primary">Submit</button>'+
        '   </form>'+
        '</div>'
})

originalAccountList1 = ""
var accountList = new Vue({
    el: '#accountRow',
    data: {
        accountList: [],
        accountList1: [],
    },
    mounted() {
        setInterval(() => {
            axios
                .get(`/api/account`)
                .then(response =>  {
                    if(!_.isEqual(originalAccountList1, response.data)) {
                        this.accountList = response.data
                        originalAccountList1 = _.cloneDeep(response.data)
                    }
                });
        }, 2000)
    },
})