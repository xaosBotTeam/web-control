package main

import (
	"github.com/xaosBotTeam/go-shared-models/apiAccountInformation"
	"github.com/xaosBotTeam/go-shared-models/dbAccountInformation"
)

type testDriver struct {
}

var accountSlice []waccount
var userList map[string]string

func init() {
	accountSlice = append(accountSlice, waccount{DbAccountInformation: dbAccountInformation.DbAccountInformation{ID: 1, FriendlyName: "acc1"}})
	accountSlice = append(accountSlice, waccount{ApiAccountInformation: apiAccountInformation.ApiAccountInformation{Sliv: true}, DbAccountInformation: dbAccountInformation.DbAccountInformation{ID: 2, FriendlyName: "acc2"}})
	accountSlice = append(accountSlice, waccount{DbAccountInformation: dbAccountInformation.DbAccountInformation{ID: 3, FriendlyName: "acc3"}})
	accountSlice = append(accountSlice, waccount{DbAccountInformation: dbAccountInformation.DbAccountInformation{ID: 4, FriendlyName: "acc4"}})
	userList = make(map[string]string)

	userList["admin"] = "admin"
}
func (driver testDriver) Authorization(login, password string) bool {
	if pass, ok := userList[login]; ok {
		if pass == password {
			return true
		}
	}

	return false
}

func (driver testDriver) GetAccountList() ([]waccount, bool) {
	return accountSlice, true
}

func (driver testDriver) GetAccountInformation(ID int) (waccount, bool) {
	for i := range accountSlice {
		if accountSlice[i].ID == ID {
			return accountSlice[i], true
		}
	}
	return waccount{}, false
}

func (driver testDriver) SetAccountInformation(ID int, account waccount) bool {
	for i := range accountSlice {
		if accountSlice[i].ID == ID {
			accountSlice[i] = account
			return true
		}
	}
	return false
}
