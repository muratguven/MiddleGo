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

func WorkflowMapToDtoList(modelList *[]models.Workflow) *[]dtos.WorkflowDto {
	var dtoList []dtos.WorkflowDto
	for _, m := range *modelList {
		var dto = &dtos.WorkflowDto{}

		dto.Id = m.Id
		dto.Name = m.Name
		dto.CreatedDate = m.CreatedDate
		dto.IsActive = m.IsActive
		dto.ModifiedDate = m.ModifiedDate
		dtoList = append(dtoList, *dto)
	}
	return &dtoList
}
