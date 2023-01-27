package Connectors

import (
	"github.com/xaosBotTeam/go-shared-models/apiAccountInformation"
	"github.com/xaosBotTeam/go-shared-models/dbAccountInformation"
)

type Waccount struct {
	apiAccountInformation.ApiAccountInformation
	dbAccountInformation.DbAccountInformation
}

type Сredentials struct {
	Login    string
	Password string
}

type Connector interface {
	Authorization(credential Сredentials) bool
	//GetUserInformation(ID int)
	GetAccountList() ([]Waccount, bool)
	GetAccountInformation(ID int) (Waccount, bool)
	SetAccountInformation(ID int, account Waccount) bool
}
