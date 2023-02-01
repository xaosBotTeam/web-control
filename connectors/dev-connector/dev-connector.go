package dev_connector

import (
	"XaocBotWebControl/connectors"
	"github.com/xaosBotTeam/go-shared-models/account"
	"github.com/xaosBotTeam/go-shared-models/task"
)

type DevConnector struct {
}

var accountSlice map[int]connectors.Waccount
var userList map[string]string

func init() {
	accountSlice = make(map[int]connectors.Waccount)
	accountSlice[0] = connectors.Waccount{account.Account{ID: 0, FriendlyName: "sasa0"}, task.Status{ArenaFarming: false}}
	accountSlice[1] = connectors.Waccount{account.Account{ID: 1, FriendlyName: "sasa1"}, task.Status{ArenaFarming: true}}
	accountSlice[2] = connectors.Waccount{account.Account{ID: 2, FriendlyName: "sasa2"}, task.Status{ArenaFarming: false}}

	userList = make(map[string]string)
	userList["admin"] = "admin1234"
}
func (connector DevConnector) Authorization(credential connectors.Сredentials) bool {
	if pass, ok := userList[credential.Login]; ok {
		if pass == credential.Password {
			return true
		}
	}

	return true
}

func (connector DevConnector) GetAccountAllInformation() (map[int]connectors.Waccount, bool) {
	return accountSlice, true
}

func (connector DevConnector) GetAccountInformation(ID int) (connectors.Waccount, bool) {
	if _, ok := accountSlice[ID]; ok {
		return accountSlice[ID], true
	}

	return connectors.Waccount{}, false
}

func (connector DevConnector) SetAccountInformation(ID int, account connectors.Waccount) bool {
	if _, ok := accountSlice[ID]; ok {
		accountSlice[ID] = account
		return true
	}
	return false
}
