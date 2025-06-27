package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	Name        string `json:"name" gorm:"column:name; type:varchar(255); not null"`
	Email       string `json:"email" gorm:"column:email; type:varchar(255); not null; unique"`
	Password    string `json:"password" gorm:"column:password; type:varchar(255)"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number; type:varchar(255)"`
	Role        string `json:"role" gorm:"column:role; type:varchar(255); default:'customer'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Order       []Order        `gorm:"foreignKey:CustomerID"`
}

func (User) TableName() string {
	return "user"
}
