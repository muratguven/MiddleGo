package models

import "time"

type User struct {
	Id           string `gorm:"primaryKey"`
	UserName     string `gorm:"index"`
	Email        *string
	Password     string
	Name         string
	LastName     string
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	ModifiedDate time.Time `gorm:"autoUpdateTime"`
	IsDeleted    bool
	IsActive     bool
}
