package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"web-control/connectors/connectors"
)

func doAuth(c *fiber.Ctx, cn connectors.Connector) (int, bool) {
	var credential = connectors.Сredentials{}

	if err := c.BodyParser(&credential); err != nil {
		return 0, false
	}
	id, auth := cn.Authorization(credential)
	if auth == false {
		return 0, false
	}

	return id, true
}

func CheckCoockes(store *session.Store, c *fiber.Ctx) bool {
	sess, err := store.Get(c)

	if sess == nil || err != nil || sess.Get("auth") == nil {
		return false
	}

	return true
}

//	@Summary		Auth
//	@ID				auth
//
//	@Tags 			Auth
//	@Produce		json
//	@Param			account	body connectors.Сredentials true "login and password"
//	@Router			/auth [post]
// @Success         200  {object}  string
// @Failure         403  {object}  string
func (controller *abstractController) Auth(c *fiber.Ctx) error {
	id, auth := doAuth(c, controller.cn)
	if !auth {
		return fiber.NewError(fiber.StatusForbidden)
	}

	sess, _ := controller.store.Get(c)
	sess.Set("auth", "ok")
	sess.Set("id", id)
	err := sess.Save()

	if err != nil {
		return fiber.NewError(fiber.StatusForbidden)
	}
	return c.JSON("ok")
}

func initAuth(app *fiber.App, controller *abstractController) {

	app.Post("/auth", controller.Auth)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/control-panel.html")
	})
	app.Get("*.html", func(c *fiber.Ctx) error {
		if CheckCoockes(controller.store, c) {
			return c.Next()
		}
		return c.Redirect("/sign-in.htm")
	})

}
