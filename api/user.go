package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type password struct {
	Password string
}

//	@Summary		Reset User Password
//	@ID				reset-user-password
//
//	@Tags 			User
//	@Produce		json
//	@Param			account	body password true "password"
//	@Router			/resetPassword [post]
// @Success         200  {object}  string
// @Failure         403  {object}  string
// @Failure         500  {object}  string
func (controller *abstractController) ResetUserPassword(c *fiber.Ctx) error {
	var pass = password{}
	c.BodyParser(&pass)

	sess, err := controller.store.Get(c)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Can not find cookie")
	}

	ID, ok := sess.Get("id").(int)

	if !ok {
		return c.Status(http.StatusBadRequest).JSON("Can not find user ID")
	}

	if !controller.cn.ResetUserPassword(ID, pass.Password) {
		return c.Status(http.StatusBadRequest).JSON("Database error")
	}
	return c.JSON("ok")
}
func initUserApi(app *fiber.App, controller *abstractController) {
	app.Post("/resetPassword", controller.ResetUserPassword)
	return
}
