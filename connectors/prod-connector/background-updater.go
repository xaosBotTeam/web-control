package prod_connector

import (
	"XaocBotWebControl/config-control-panel"
	"XaocBotWebControl/connectors"
	"encoding/json"
	"github.com/xaosBotTeam/go-shared-models/account"
	"github.com/xaosBotTeam/go-shared-models/config"
	"github.com/xaosBotTeam/go-shared-models/status"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func parser(body []byte) (string, bool) {
	var objmap map[string]json.RawMessage
	err := json.Unmarshal(body, &objmap)
	if err != nil {
		log.Println(err)
		return "", true
	}

	var strJson string
	err = json.Unmarshal(objmap["data"], &strJson)
	if err != nil {
		log.Println(err)
		return "", true
	}

	return strJson, false
}

func updateAccountStatus() (map[string]status.Status, bool) {
	resp, err := http.Get(config_control_panel.Configuration.URL_GO_BOT + "/status")
	if err != nil {
		log.Println(err)
		return nil, true
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, true
	}

	if strAccountStatus, err_ := parser(body); !err_ {
		AccountStatus := map[string]status.Status{}

		err = json.Unmarshal([]byte(strAccountStatus), &AccountStatus)
		if err != nil {
			log.Println(err)
			return nil, true
		}

		return AccountStatus, false
	}
	return nil, true

}

func updateAccountInformation() (map[string]account.Account, bool) {
	resp, err := http.Get(config_control_panel.Configuration.URL_GO_BOT + "/account")
	if err != nil {
		log.Println(err)
		return nil, true
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, true
	}

	if strAccountInformation, err_ := parser(body); !err_ {
		AccountInformation := map[string]account.Account{}

		err = json.Unmarshal([]byte(strAccountInformation), &AccountInformation)
		if err != nil {
			log.Println(err)
			return AccountInformation, true
		}

		return AccountInformation, false
	}
	return nil, true

}

func updateAccountConfig() (map[string]config.Config, bool) {
	resp, err := http.Get(config_control_panel.Configuration.URL_GO_BOT + "/config")
	if err != nil {
		log.Println(err)
		return nil, true
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, true
	}

	if strAccountConfig, err_ := parser(body); !err_ {
		AccountConfig := map[string]config.Config{}

		err = json.Unmarshal([]byte(strAccountConfig), &AccountConfig)
		if err != nil {
			log.Println(err)
			return nil, true
		}

		return AccountConfig, false
	}
	return nil, true

}

func backgroundUpdater() {
	for {
		accounts := map[string]connectors.Waccount{}

		status, err := updateAccountStatus()

		if err {
			time.Sleep(5 * time.Second)
			continue
		}

		for k, v := range status {
			accounts[k] = connectors.Waccount{Status: v}
		}

		inf, err := updateAccountInformation()

		if err {
			time.Sleep(5 * time.Second)
			continue
		}
		for key, inf := range inf {
			if _, err := accounts[key]; err {
				account := accounts[key]
				account.AddAccount(inf)
				accounts[key] = account
			}
		}

		for i := 0; i < 30; i++ {
			if conf, err := updateAccountConfig(); err == false {
				for key, conf := range conf {
					if _, err := accounts[key]; err {
						account := accounts[key]
						account.AddConfig(conf)
						accounts[key] = account
					}
				}

				resAccount := map[int]connectors.Waccount{}

				for key, acc := range accounts {
					intKey, err := strconv.Atoi(key)
					if err != nil {
						break
					}

					resAccount[intKey] = acc
				}
				accountStorage.Lock()
				accountStorage.accountMap = resAccount
				accountStorage.Unlock()
			}
			time.Sleep(1 * time.Second)
		}

	}
}
