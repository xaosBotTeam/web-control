package Connectors

import (
	"github.com/xaosBotTeam/go-shared-models/apiAccountInformation"
	"github.com/xaosBotTeam/go-shared-models/dbAccountInformation"
)

type TestDriver struct {
}

var accountSlice []Waccount
var userList map[string]string

func init() {
	accountSlice = append(accountSlice, Waccount{DbAccountInformation: dbAccountInformation.DbAccountInformation{ID: 1, FriendlyName: "acc1"}})
	accountSlice = append(accountSlice, Waccount{ApiAccountInformation: apiAccountInformation.ApiAccountInformation{Sliv: true}, DbAccountInformation: dbAccountInformation.DbAccountInformation{ID: 2, FriendlyName: "acc2"}})
	accountSlice = append(accountSlice, Waccount{DbAccountInformation: dbAccountInformation.DbAccountInformation{ID: 3, FriendlyName: "acc3"}})
	accountSlice = append(accountSlice, Waccount{DbAccountInformation: dbAccountInformation.DbAccountInformation{ID: 4, FriendlyName: "acc4"}})
	userList = make(map[string]string)

	userList["admin"] = "admin"
}
func (driver TestDriver) Authorization(credential Сredentials) bool {

	if pass, ok := userList[credential.Login]; ok {
		if pass == credential.Password {
			return true
		}
	}

	return false
}

func (driver TestDriver) GetAccountList() ([]Waccount, bool) {
	return accountSlice, true
}

func (driver TestDriver) GetAccountInformation(ID int) (Waccount, bool) {
	for i := range accountSlice {
		if accountSlice[i].ID == ID {
			return accountSlice[i], true
		}
	}
	return Waccount{}, false
}

func (driver TestDriver) SetAccountInformation(ID int, account Waccount) bool {
	for i := range accountSlice {
		if accountSlice[i].ID == ID {
			accountSlice[i] = account
			return true
		}
	}
	return false
}
