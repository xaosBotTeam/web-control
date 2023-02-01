Vue.component('account-item', {
    methods: {
        formSubmit(e) {
            console.log(this)
            e.preventDefault();
            axios.put(`/api/config/${this.$vnode.key}`, {
                arena_farming: this.account.arena_farming,
                arena_use_energy_cans: this.account.arena_use_energy_cans,
                travelling: this.account.travelling,
            }).catch(this.account.arena_farming=originalAccountList1[this.$vnode.key].arena_farming);
        }
    },

    props: ['account','pID'],

    template:
        '<div class="col">' +
        '   <form @submit.prevent="formSubmit">'+
        '       <a target="_blank" rel="noopener noreferrer" :href="account.url"><h4>{{account.friendly_name}}</h4></a>' +
        '            <div class="form-check form-switch">' +
        '                <input class="form-check-input" type="checkbox" role="switch" v-model="account.arena_farming" id="check_arena_farming">' +
        '                <label class="form-check-label" for="check_arena_farming">Слив</label>' +
        '            </div>' +
        '            <div class="form-check form-switch">' +
        '                <input class="form-check-input" type="checkbox" role="switch" v-model="account.arena_use_energy_cans" id="arena_use_energy_cans">' +
        '                <label class="form-check-label" for="arena_use_energy_cans">Использовать банки</label>' +
        '            </div>' +
        '            <div class="form-check form-switch">' +
        '                <input class="form-check-input" type="checkbox" role="switch" v-model="account.travelling" id="travelling">' +
        '                <label class="form-check-label" for="arena_use_energy_cans">Путешествия</label>' +
        '            </div>' +
        '            <details>' +
        '            <summary>Details</summary>' +
        '               <div>game_id - {{account.game_id}}</div>' +
        '               <div>energy_limit - {{account.energy_limit}}</div>' +
        '            </details>'+
        '       <button class="btn btn-primary">Submit</button>'+
        '   </form>'+
        '</div>'
})

originalAccountList1 = ""
var accountList = new Vue({
    el: '#accountRow',
    data: {
        accountList: [],
    },
    mounted() {
        axios
            .get(`/api/account`)
            .then(response =>  {
                if(!_.isEqual(originalAccountList1, response.data)) {
                    this.accountList = response.data
                    originalAccountList1 = _.cloneDeep(response.data)
                }
            });

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


var addAccount = new Vue({
    el: '#addAccountForm',
    data: {
        url: ''
    },
    methods: {
        addAccountFormSubmit(e) {
            console.log(this)
            axios.post(`/api/account`, {
                url: this.url,
            })
            this.url=''
            document.getElementById('close').click();
        }
    },
})