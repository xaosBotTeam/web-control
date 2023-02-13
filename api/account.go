package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/xaosBotTeam/go-shared-models/config"
	"net/http"
	"strconv"
	"web-control/connectors/connectors"
)

func Account(app *fiber.App, store *session.Store, cn connectors.Connector) {
	api := app.Group("/api")
	api.Use(func(c *fiber.Ctx) error {
		if CheckCoockes(store, c) {
			return c.Next()
		}
		return c.Status(http.StatusForbidden).JSON("Can't parse account id")
	})

	api.Get("/account/:accountID", func(c *fiber.Ctx) error {
		ids := c.Params("accountID")

		id, err := strconv.Atoi(ids)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON("Can't parse account id")
		}

		account_, ok := cn.GetAccountInformation(id)
		if !ok {
			return fiber.NewError(fiber.StatusBadGateway)
		}

		return c.JSON(account_)
	})

	api.Delete("/account/:accountID", func(c *fiber.Ctx) error {
		ids := c.Params("accountID")

		id, err := strconv.Atoi(ids)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON("Can't parse account id")
		}

		cn.DeleteAccount(id)

		return c.JSON("ok")
	})

	api.Put("/config/:accountID", func(c *fiber.Ctx) error {

		ids := c.Params("accountID")

		id, err := strconv.Atoi(ids)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON("Can't parse account id")
		}

		var config = config.Config{}
		account_, ok := cn.GetAccountInformation(id)

		if !ok {
			return c.Status(fiber.StatusNotFound).JSON("Can't find account with such id")
		}

		if err := c.BodyParser(&config); err != nil {
			return c.Status(http.StatusBadRequest).JSON("Can't parse body")
		}

		account_.AddConfig(config)

		cn.SetAccountInformation(id, account_)
		return c.JSON(account_)
	})

	api.Get("/account", func(c *fiber.Ctx) error {
		AccountList, ok := cn.GetAccountAllInformation()
		if !ok {
			return fiber.NewError(fiber.StatusBadGateway)
		}
		return c.JSON(AccountList)
	})

	api.Post("/account", func(c *fiber.Ctx) error {

		type Url struct {
			Url string
		}
		var url Url
		if err := c.BodyParser(&url); err != nil {
			return c.Status(http.StatusBadRequest).JSON("Can't parse body")
		}
		cn.CreateAccount(url.Url)
		return c.JSON(url)
	})

}
