package prod_connector

import (
	"encoding/json"
	"github.com/xaosBotTeam/go-shared-models/account"
	"github.com/xaosBotTeam/go-shared-models/config"
	"strconv"
	"sync"
	configControlPanel "web-control/connectors/config-control-panel"
	"web-control/connectors/connectors"
	"web-control/connectors/connectors/prod-connector/postgres-connector"
	sending_queue "web-control/connectors/connectors/prod-connector/sending-queue"
)

type ProdConnector struct {
}

var accountStorage AccountStorage

type AccountStorage struct {
	accountMap map[int]connectors.FullAccount
	sync.RWMutex
}

func init() {
	accountStorage.accountMap = make(map[int]connectors.FullAccount)
	go backgroundUpdater()
}

func (connector ProdConnector) Authorization(credential connectors.Сredentials) (int, bool) {
	id, resp := postgres_connector.AuthUser(credential.Login, credential.Password)
	return id, resp
}

func (connector ProdConnector) GetAccountAllInformation() (map[int]connectors.FullAccount, bool) {
	accountStorage.RLock()
	defer accountStorage.RUnlock()
	return accountStorage.accountMap, true
}

func (connector ProdConnector) GetAccountInformation(ID int) (connectors.FullAccount, bool) {
	accountStorage.RLock()
	defer accountStorage.RUnlock()
	if _, ok := accountStorage.accountMap[ID]; ok {
		return accountStorage.accountMap[ID], true
	}

	return connectors.FullAccount{}, false
}

func (connector ProdConnector) SetAccountInformation(ID int, account connectors.FullAccount) bool {
	accountStorage.RLock()
	defer accountStorage.RUnlock()
	if _, ok := accountStorage.accountMap[ID]; ok {
		b, err := json.Marshal(config.Config{ArenaFarming: account.ArenaFarming, ArenaUseEnergyCans: account.ArenaUseEnergyCans, Travelling: account.Travelling})
		if err != nil {
			return false
		}
		queue := sending_queue.Queue_{Url: configControlPanel.GetBotURl() + "/config/" + strconv.Itoa(ID), Methods: "PUT", Value: b}
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

	queue := sending_queue.Queue_{Url: configControlPanel.GetBotURl() + "/account/", Methods: "POST", Value: b}
	sending_queue.Channel <- queue

	queue = sending_queue.Queue_{Url: configControlPanel.GetBotURl() + "/refresh/", Methods: "PATCH", Value: nil}
	sending_queue.Channel <- queue
}

func (connector ProdConnector) ResetUserPassword(ID int, password string) bool {
	return postgres_connector.ChangePassword(ID, password)
}
