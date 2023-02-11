package config_control_panel

import (
	"log"
	"os"
	"time"
)

type Configurations struct {
	ready  bool
	botUrl string
}

var configuration Configurations

func init() {
	GO_BOT_URL := os.Getenv("GO_BOT_URL")
	if GO_BOT_URL == "" {
		log.Println("WARNING: GO_BOT_URL is empty. Set default value http://127.0.0.1:5504")
		configuration.botUrl = "http://127.0.0.1:5504"
	} else {
		log.Println("Set GO_BOT_URL ", GO_BOT_URL)
		configuration.botUrl = GO_BOT_URL
	}
	configuration.ready = true
}

func GetBotURl() string {
	for !configuration.ready {
		time.Sleep(1 * time.Second)
	}
	return configuration.botUrl
}
