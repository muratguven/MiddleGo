package api

import "github.com/gofiber/fiber/v2"

func (app *Application) Routes(fApp *fiber.App) {
	fApp.Post("/AddWorkflow", app.AddWorkflow)
}
