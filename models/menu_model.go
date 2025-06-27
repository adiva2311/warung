package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID         uint    `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	Name       string  `json:"name" gorm:"column:name; type:varchar(255)"`
	Price      float32 `json:"price" gorm:"column:price; type:float"`
	UserID     uint    `json:"user_id" gorm:"column:user_id; type:int"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	CategoryID uint           `json:"category_id" gorm:"column:category_id"`
	Category   Category
}

func (Menu) TableName() string {
	return "menu"
}
