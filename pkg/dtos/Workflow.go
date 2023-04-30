package dtos

type WorkflowDto struct {
	BaseDto
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}
