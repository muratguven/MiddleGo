package models

// It is test domain :)
type User struct {
	BaseEntity
	UserName string  `gorm:"index;column:UserName;not null"`
	Email    *string `gorm:"column:Email;not null"`
	Password string  `gorm:"column:Password"`
	Name     string  `gorm:"column:Name"`
	LastName string  `gorm:"column:LastName"`
	IsActive bool    `gorm:"column:IsActive;not null"`
}
