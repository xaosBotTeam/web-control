package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xaosBotTeam/go-shared-models/config"
	"net/http"
	"strconv"
)

//	@Summary		Get account by id
//	@ID				get-account
//
// @Tags 			Account
// @Accept 			json
// @Router			/account/{id} [get]
// @Success         200  {object} connectors.FullAccount
// @Failure         403  {object}  string
// @Failure         500  {object}  string
// @Param			id	path int true "account id"
func (controller *abstractController) GetAccount(c *fiber.Ctx) error {
	ids := c.Params("accountID")

	id, err := strconv.Atoi(ids)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Can't parse account id")
	}

	account, ok := controller.cn.GetAccountInformation(id)
	if !ok {
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return c.JSON(account)
}

//	@Summary		Delete game account by id
//	@ID				delete-account-by-id
//
//	@Tags 			Account
//	@Produce		json
//	@Param			id	path int true "account id"
//	@Router			/account/{id} [delete]
// @Success         200  {object}  string
// @Failure         403  {object}  string
// @Failure         500  {object}  string
func (controller *abstractController) DeleteAccount(c *fiber.Ctx) error {
	ids := c.Params("accountID")

	id, err := strconv.Atoi(ids)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Can't parse account id")
	}

	controller.cn.DeleteAccount(id)

	return c.JSON("ok")
}

//	@Summary		Update account by id
//	@ID				update-account-config
// @Tags 			Account
// @Accept 			json
// @Router			/account/{id} [put]
//	@Param			id	path int true "account id"
// @Success         200  {object}  connectors.FullAccount
// @Failure         403  {object}  string
// @Failure         500  {object}  string
func (controller *abstractController) UpdateAccount(c *fiber.Ctx) error {

	ids := c.Params("accountID")

	id, err := strconv.Atoi(ids)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON("Can't parse account id")
	}

	var config = config.Config{}
	account_, ok := controller.cn.GetAccountInformation(id)

	if !ok {
		return c.Status(fiber.StatusNotFound).JSON("Can't find account with such id")
	}

	if err := c.BodyParser(&config); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Can't parse body")
	}

	account_.AddConfig(config)

	controller.cn.SetAccountInformation(id, account_)
	return c.JSON(account_)
}

//	@Summary		Get all game accounts
//	@ID				get-all-accounts
//
//	@Tags 			Account
//	@Produce		json
//	@Router			/account/ [get]
// @Success         200  {array} connectors.FullAccount
// @Failure         403  {object}  string
// @Failure         500  {object}  string
func (controller *abstractController) GetAllAccounts(c *fiber.Ctx) error {
	AccountList, ok := controller.cn.GetAccountAllInformation()
	if !ok {
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return c.JSON(AccountList)
}

//	@Summary		Add new game account
//	@ID				add-new-game-account
//
//	@Tags 			Account
//	@Produce		json
//	@Param			account	body account.Account true "account url"
//	@Router			/account/ [post]
// @Success         200  {object}  account.Account
// @Failure         403  {object}  string
// @Failure         500  {object}  string
func (controller *abstractController) CreateAccount(c *fiber.Ctx) error {

	type Url struct {
		Url string
	}
	var url Url
	if err := c.BodyParser(&url); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Can't parse body")
	}
	controller.cn.CreateAccount(url.Url)
	return c.JSON(url)
}

func initAccountApi(app *fiber.App, controller *abstractController) {
	api := app.Group("/api")
	api.Use(func(c *fiber.Ctx) error {
		if CheckCoockes(controller.store, c) {
			return c.Next()
		}
		return c.Status(http.StatusForbidden).JSON("Can't parse account id")
	})

	api.Get("/account/:accountID", controller.GetAccount)

	api.Delete("/account/:accountID", controller.DeleteAccount)

	api.Put("/config/:accountID", controller.UpdateAccount)

	api.Get("/account", controller.GetAllAccounts)

	api.Post("/account", controller.CreateAccount)

}
