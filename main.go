package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"web-control/connectors/api"
	"web-control/connectors/connectors"
	prod_connector "web-control/connectors/connectors/prod-connector"
)

func main() {
	//var cn connectors.Connector = dev_connector.DevConnector{}
	var cn connectors.Connector = prod_connector.ProdConnector{}

	app := fiber.New()
	store := session.New()
	api.Account(app, store, cn)
	api.Auth(app, store, cn)
	api.User(app, store, cn)

	app.Static("/", "./static")
	log.Println(app.Listen(":3000"))
}
