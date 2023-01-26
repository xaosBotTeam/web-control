package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func configWebServer(cn Connector) {
	app := fiber.New()

	api := app.Group("/api")

	api.Get("/account", func(c *fiber.Ctx) error {
		AccountList, ok := cn.GetAccountList()
		if !ok {
			return fiber.NewError(fiber.StatusBadGateway)
		}
		return c.JSON(AccountList)
	})

	api.Get("/account/:accountID", func(c *fiber.Ctx) error {

		ids := c.Params("accountID")

		id, err := strconv.Atoi(ids)
		if err != nil {
			return fiber.NewError(fiber.StatusBadGateway)
		}

		account_, ok := cn.GetAccountInformation(id)
		if !ok {
			return fiber.NewError(fiber.StatusBadGateway)
		}

		return c.JSON(account_)
	})

	api.Post("/account/:accountID", func(c *fiber.Ctx) error {

		ids := c.Params("accountID")

		id, err := strconv.Atoi(ids)
		if err != nil {
			return fiber.NewError(fiber.StatusBadGateway)
		}

		var account = waccount{}
		account_, ok := cn.GetAccountInformation(id)

		if !ok {
			return fiber.NewError(fiber.StatusBadGateway)
		}

		if err := c.BodyParser(&account); err != nil {
			return fiber.NewError(fiber.StatusBadGateway)
		}

		//TODO: fix this
		if account.FriendlyName != "" {
			account_.FriendlyName = account.FriendlyName
		}

		account_.Owner = account.Owner

		account_.Sliv = account.Sliv

		cn.SetAccountInformation(id, account_)
		return c.JSON(account_)
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/controlPanel.html")
	})
	app.Static("/", "./static")
	log.Println(app.Listen(":3000"))
}
