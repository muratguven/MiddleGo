package utils

import (
	"middle/pkg/dtos"
	"middle/pkg/models"
	"time"
)

func WorkflowMapToModel(dto *dtos.WorkflowDto) *models.Workflow {
	var model = &models.Workflow{}

	model.Name = dto.Name
	model.IsActive = dto.IsActive
	model.IsDeleted = false
	model.CreatedDate = time.Now()
	return model
}
