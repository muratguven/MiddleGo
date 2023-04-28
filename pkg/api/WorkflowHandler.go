package api

import (
	"context"
	"log"
	"middle/pkg/dtos"
	"middle/pkg/infrastructures/repository"
	"middle/pkg/models"

	"middle/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func (app *Application) AddWorkflow(c *fiber.Ctx) error {
	ctx := context.Background()
	workflow := &dtos.WorkflowDto{}
	if err := c.BodyParser(&workflow); err != nil {
		return c.SendStatus(403)
	}
	repo := repository.NewRepository[models.Workflow](app.Db)
	err := repo.Insert(utils.WorkflowMapToModel(workflow), ctx)

	if err != nil {
		log.Fatal(err)
		return c.SendStatus(500)
	}

	return c.SendStatus(200)

}
