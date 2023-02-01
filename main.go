package main

import (
	"XaocBotWebControl/api"
	"XaocBotWebControl/connectors"
	prod_connector "XaocBotWebControl/connectors/prod-connector"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
)

func main() {
	//var cn connectors.Connector = dev_connector.DevConnector{}
	var cn connectors.Connector = prod_connector.ProdConnector{}

	app := fiber.New()
	store := session.New()
	api.Account(app, cn)
	api.Auth(app, store, cn)

	app.Static("/", "./static")
	log.Println(app.Listen(":3000"))
}
