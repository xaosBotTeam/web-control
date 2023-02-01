package config_control_panel

import (
	"log"
	"os"
)

type Configurations struct {
	GO_BOT_URL string
}

var Configuration Configurations

func init() {
	GO_BOT_URL := os.Getenv("GO_BOT_URL")
	if GO_BOT_URL == "" {
		log.Println("WARNING: GO_BOT_URL is empty. Set default value http://127.0.0.1:5504")
		Configuration.GO_BOT_URL = "http://127.0.0.1:5504"
	} else {
		log.Println("Set GO_BOT_URL ", GO_BOT_URL)
		Configuration.GO_BOT_URL = GO_BOT_URL
	}

}
