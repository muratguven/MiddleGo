package api

import (
	"context"
	"log"
	"middle/pkg/domains"
	"middle/pkg/dtos"
	"middle/pkg/infrastructures/repository"

	"middle/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// Handlers are like controllers :)
func (app *Application) Add(c *fiber.Ctx) error {
	ctx := context.Background()
	workflow := &dtos.WorkflowDto{}
	if err := c.BodyParser(&workflow); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}
	repo := repository.NewRepository[domains.Workflow](app.Db)
	err := repo.Insert(utils.WorkflowMapToModel(workflow), ctx)

	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occured while adding workflow"))
	}

	return c.Status(fiber.StatusOK).JSON(utils.StatusOK("Workflow added successfully."))

}

func (app *Application) GetAll(c *fiber.Ctx) error {
	ctx := context.Background()
	repo := repository.NewRepository[domains.Workflow](app.Db)
	dataList, err := repo.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occured while getting workflows"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK(utils.WorkflowMapToDtoList(dataList)))
}
