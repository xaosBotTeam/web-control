package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/swaggo/fiber-swagger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"log"
	"web-control/connectors/api"
	"web-control/connectors/connectors"
	prod_connector "web-control/connectors/connectors/prod-connector"
	_ "web-control/connectors/docs"
)

// @title API for xaos bot web control
// @version 1.0.0
// @description API for xaos bot web control
// @BasePath /api/

func main() {
	var cn connectors.Connector = prod_connector.ProdConnector{}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
		AppName:               "Web-control",
	})
	store := session.New()
	api.InitRoutes(app, store, cn)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Static("/", "./static")
	log.Fatalln(app.Listen(":3000"))
}
