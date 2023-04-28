package models

type Workflow struct {
	BaseEntity
	Name     string `gorm:"column:Name;not null"`
	IsActive bool   `gorm:"column:IsActive;not null"`
}

func (Workflow) TableName() string {
	return "WorkFlows"
}
