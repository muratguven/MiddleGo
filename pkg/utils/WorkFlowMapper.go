package utils

import (
	"middle/pkg/domains"
	"middle/pkg/dtos"
	"time"
)

func WorkflowMapToModel(dto *dtos.WorkflowDto) *domains.Workflow {
	var model = &domains.Workflow{}

	model.Name = dto.Name
	model.IsActive = dto.IsActive
	model.IsDeleted = false
	model.CreatedDate = time.Now()
	return model
}

func WorkflowMapToDtoList(modelList *[]domains.Workflow) *[]dtos.WorkflowDto {
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
