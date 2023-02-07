package connectors

import (
	"github.com/xaosBotTeam/go-shared-models/account"
	"github.com/xaosBotTeam/go-shared-models/config"
	"github.com/xaosBotTeam/go-shared-models/status"
)

type Waccount struct {
	account.Account
	config.Config
	status.Status
}

func (wacc *Waccount) AddAccount(acc account.Account) {
	wacc.URL = acc.URL
	wacc.Owner = acc.Owner
}

func (wacc *Waccount) AddConfig(conf config.Config) {
	wacc.ArenaUseEnergyCans = conf.ArenaUseEnergyCans
	wacc.ArenaFarming = conf.ArenaFarming
	wacc.Travelling = conf.Travelling
}

func (wacc *Waccount) AddStatus(status status.Status) {
	wacc.FriendlyName = status.FriendlyName
	wacc.GameID = status.GameID
	wacc.EnergyLimit = status.EnergyLimit
}

type Сredentials struct {
	Login    string
	Password string
}

type Connector interface {
	Authorization(credential Сredentials) (int, bool)

	GetAccountAllInformation() (map[int]Waccount, bool)
	GetAccountInformation(ID int) (Waccount, bool)

	SetAccountInformation(ID int, account Waccount) bool
	CreateAccount(url string)

	ResetUserPassword(ID int, password string) bool
}
