Vue.component('account-item', {
    methods: {
        formSubmit(e) {
            e.preventDefault();
            axios.put(`/api/config/${this.$vnode.key}`, {
                arena_farming: this.account.arena_farming,
                arena_use_energy_cans: this.account.arena_use_energy_cans,
                travelling: this.account.travelling,
                open_chests: this.account.open_chests,
            }).catch(this.account.arena_farming = originalAccountList1[this.$vnode.key].arena_farming);
        },
        deleteAccount(e) {
            e.preventDefault();
            axios.delete(`/api/account/${this.$vnode.key}`, {
            }).catch(this.account.arena_farming = originalAccountList1[this.$vnode.key].arena_farming);
        },
    },

    props: ['account', 'pID'],


    template:

`
<div class="col py-3">
  <form>
    <div class="row">
      <div class="col informationIcon">
       <a target="_blank" rel="noopener noreferrer" :href="account.url">
        <h4>{{account.friendly_name}}</h4>
       </a>
        <div class="dropdown px-2">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" id="dropdownMenuButton1" data-bs-toggle="dropdown" aria-expanded="false" fill="currentColor" className="bi bi-info-circle" viewBox="0 0 16 16">
            <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
            <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z" />
          </svg>
          <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
            <li><div class="dropdown-item" @click="deleteAccount">Удалить аккаунт</div></li>
          </ul>
        </div>
      </div>
    </div>
    <div class="form-check form-switch">
      <input class="form-check-input" type="checkbox" role="switch" v-model="account.arena_farming" @change="formSubmit">
      <label class="form-check-label">Слив</label>
    </div>
    <div class="form-check form-switch">
      <input class="form-check-input" type="checkbox" role="switch" v-model="account.arena_use_energy_cans" @change="formSubmit">
      <label class="form-check-label">Использовать банки</label>
    </div>
    <div class="form-check form-switch">
      <input class="form-check-input" type="checkbox" role="switch" v-model="account.travelling" @change="formSubmit">
      <label class="form-check-label">Путешествия</label>
    </div>
    <div class="form-check form-switch">
      <input class="form-check-input" type="checkbox" role="switch" v-model="account.open_chests" @change="formSubmit">
      <label class="form-check-label">Открывать кейсы</label>
    </div>
    <details>
      <summary>Details</summary>
      <div>game_id - {{account.game_id}}</div>
      <div>energy_limit - {{account.energy_limit}}</div>
    </details>
  </form>
</div>
`
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
            .then(response => {
                if (!_.isEqual(originalAccountList1, response.data)) {
                    this.accountList = response.data
                    originalAccountList1 = _.cloneDeep(response.data)
                }
            });

        setInterval(() => {
            axios
                .get(`/api/account`)
                .then(response => {
                    if (!_.isEqual(originalAccountList1, response.data)) {
                        this.accountList = response.data
                        originalAccountList1 = _.cloneDeep(response.data)
                    }
                })
                .catch(error => {
                    if (error.response.status == 403) {
                        window.location.href = "sign-in.js.html"
                    }
                    console.log(error)
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
            this.url = ''
            document.getElementById('close').click();
        }
    },
})