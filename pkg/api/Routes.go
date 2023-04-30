package api

import "github.com/gofiber/fiber/v2"

func (app *Application) Routes(fApp *fiber.App) {
	fApp.Post("api/workflow/add", app.Add)
	fApp.Get("api/workflow/all", app.GetAll)
}
