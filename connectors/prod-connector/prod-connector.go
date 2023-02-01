package prod_connector

import (
	"encoding/json"
	"github.com/xaosBotTeam/go-shared-models/account"
	"github.com/xaosBotTeam/go-shared-models/config"
	"strconv"
	"sync"
	config_control_panel "web-control/connectors/config-control-panel"
	"web-control/connectors/connectors"
	sending_queue "web-control/connectors/connectors/prod-connector/sending-queue"
)

type ProdConnector struct {
}

var accountStorage AccountStorage
var userList map[string]string

type AccountStorage struct {
	accountMap map[int]connectors.Waccount
	sync.Mutex
}

func init() {
	accountStorage.accountMap = make(map[int]connectors.Waccount)

	userList = make(map[string]string)
	userList["admin"] = "admin1234"

	go backgroundUpdater()
}

func (connector ProdConnector) Authorization(credential connectors.Сredentials) bool {
	if pass, ok := userList[credential.Login]; ok {
		if pass == credential.Password {
			return true
		}
	}

	return true
}

func (connector ProdConnector) GetAccountAllInformation() (map[int]connectors.Waccount, bool) {
	accountStorage.Lock()
	defer accountStorage.Unlock()
	return accountStorage.accountMap, true
}

func (connector ProdConnector) GetAccountInformation(ID int) (connectors.Waccount, bool) {
	accountStorage.Lock()
	defer accountStorage.Unlock()
	if _, ok := accountStorage.accountMap[ID]; ok {
		return accountStorage.accountMap[ID], true
	}

	return connectors.Waccount{}, false
}

func (connector ProdConnector) SetAccountInformation(ID int, account connectors.Waccount) bool {
	accountStorage.Lock()
	defer accountStorage.Unlock()
	if _, ok := accountStorage.accountMap[ID]; ok {
		b, err := json.Marshal(config.Config{ArenaFarming: account.ArenaFarming, ArenaUseEnergyCans: account.ArenaUseEnergyCans, Travelling: account.Travelling})
		if err != nil {
			return false
		}
		queue := sending_queue.Queue_{Url: config_control_panel.Configuration.GO_BOT_URL + "/config/" + strconv.Itoa(ID), Methods: "PUT", Value: b}
		sending_queue.Channel <- queue
		return true
	}
	return false
}

func (connector ProdConnector) CreateAccount(url string) {
	b, err := json.Marshal(account.Account{Owner: 0, URL: url})
	if err != nil {
		return
	}

	queue := sending_queue.Queue_{Url: config_control_panel.Configuration.GO_BOT_URL + "/account/", Methods: "POST", Value: b}
	sending_queue.Channel <- queue

	queue = sending_queue.Queue_{Url: config_control_panel.Configuration.GO_BOT_URL + "/refresh/", Methods: "PATCH", Value: nil}
	sending_queue.Channel <- queue
}
