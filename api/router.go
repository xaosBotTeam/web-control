package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"web-control/connectors/connectors"
)

type abstractController struct {
	store *session.Store
	cn    connectors.Connector
}

func InitRoutes(app *fiber.App, store *session.Store, cn connectors.Connector) {
	var controller abstractController
	controller.store = store
	controller.cn = cn

	initAccountApi(app, &controller)
	initUserApi(app, &controller)
	initAuth(app, &controller)

}
