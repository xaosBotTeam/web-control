package config_control_panel

type Configurations struct {
	URL_GO_BOT string
}

var Configuration Configurations

func init() {
	Configuration.URL_GO_BOT = "http://127.0.0.1:5504"
}
