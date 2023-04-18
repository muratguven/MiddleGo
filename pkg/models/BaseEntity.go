package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	Id           uuid.UUID `gorm:"column:Id;primary_key;"`
	IsDeleted    bool      `gorm:"column:IsDeleted;default:false"`
	CreatedDate  time.Time `gorm:"column:CreatedDate;autoCreateTime"`
	ModifiedDate time.Time `gorm:"column:ModifiedDate;autoUpdateTime"`
}

//Hook function
func (baseEntity *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {

	baseEntity.Id = uuid.NewV4()

	return
}
