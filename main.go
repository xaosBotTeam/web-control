package main

import (
	"github.com/xaosBotTeam/go-shared-models/apiAccountInformation"
	"github.com/xaosBotTeam/go-shared-models/dbAccountInformation"
)

type waccount struct {
	apiAccountInformation.ApiAccountInformation
	dbAccountInformation.DbAccountInformation
}

type Connector interface {
	Authorization(login string, password string) bool
	//GetUserInformation(ID int)
	GetAccountList() ([]waccount, bool)
	GetAccountInformation(ID int) (waccount, bool)

	SetAccountInformation(ID int, account waccount) bool
}

func main() {
	var cn Connector = testDriver{}
	configWebServer(cn)
}
