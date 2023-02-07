package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
	"web-control/connectors/connectors"
)

func User(app *fiber.App, store *session.Store, cn connectors.Connector) {

	app.Post("/resetPassword", func(c *fiber.Ctx) error {
		type newPassword struct {
			password string
		}
		var pass = newPassword{}
		c.BodyParser(&pass)

		sess, err := store.Get(c)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON("Can not find cookie")
		}

		ID, ok := sess.Get("id").(int)

		if !ok {
			return c.Status(http.StatusBadRequest).JSON("Can not find user ID")
		}

		if !cn.ResetUserPassword(ID, pass.password) {
			return c.Status(http.StatusBadRequest).JSON("Database error")
		}
		return c.JSON("ok")
	})
	return
}
