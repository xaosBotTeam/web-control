package connectors

import (
	"github.com/xaosBotTeam/go-shared-models/account"
	"github.com/xaosBotTeam/go-shared-models/config"
	"github.com/xaosBotTeam/go-shared-models/status"
)

type FullAccount struct {
	account.Account
	config.Config
	status.Status
}

func (fullAcc *FullAccount) AddAccount(acc account.Account) {
	fullAcc.URL = acc.URL
	fullAcc.Owner = acc.Owner
}

func (fullAcc *FullAccount) AddConfig(conf config.Config) {
	fullAcc.ArenaUseEnergyCans = conf.ArenaUseEnergyCans
	fullAcc.ArenaFarming = conf.ArenaFarming
	fullAcc.Travelling = conf.Travelling
}

func (fullAcc *FullAccount) AddStatus(status status.Status) {
	fullAcc.FriendlyName = status.FriendlyName
	fullAcc.GameID = status.GameID
	fullAcc.EnergyLimit = status.EnergyLimit
}

type Сredentials struct {
	Login    string
	Password string
}

type Connector interface {
	Authorization(credential Сredentials) (int, bool)

	GetAccountAllInformation() (map[int]FullAccount, bool)
	GetAccountInformation(ID int) (FullAccount, bool)

	SetAccountInformation(ID int, account FullAccount) bool
	CreateAccount(url string)

	ResetUserPassword(ID int, password string) bool
}
