package Api

import (
	"XaocBotWebControl/Connectors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func doAuth(c *fiber.Ctx, cn Connectors.Connector) bool {
	var credential = Connectors.Сredentials{}

	if err := c.BodyParser(&credential); err != nil {
		return false
	}
	if cn.Authorization(credential) == false {
		return false
	}

	return true
}

func CheckCoockes(store *session.Store, c *fiber.Ctx) bool {
	sess, err := store.Get(c)

	if sess == nil || err != nil || sess.Get("auth") == nil {
		return false
	}

	return true
}

func Auth(app *fiber.App, store *session.Store, cn Connectors.Connector) {
	app.Post("/auth", func(c *fiber.Ctx) error {
		if !doAuth(c, cn) {
			return fiber.NewError(fiber.StatusForbidden)
		}

		sess, _ := store.Get(c)
		sess.Set("auth", "ok")
		err := sess.Save()

		if err != nil {
			return fiber.NewError(fiber.StatusForbidden)
		}
		return c.JSON("ok")

	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/control-panel.html")
	})
	app.Get("*.html", func(c *fiber.Ctx) error {

		if CheckCoockes(store, c) {
			return c.Next()
		}
		return c.Redirect("/sign-in.htm")
	})

}
