package main

import (
	"XaocBotWebControl/Api"
	"XaocBotWebControl/Connectors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
)

func main() {
	var cn Connectors.Connector = Connectors.TestDriver{}
	app := fiber.New()
	store := session.New()
	Api.Account(app, cn)
	Api.Auth(app, store, cn)

	app.Static("/", "./static")
	log.Println(app.Listen(":3000"))
}
