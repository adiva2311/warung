package models

import "time"

type Order struct {
	ID         uint `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	CustomerID uint `json:"customer_id"`
	Customer   User `gorm:"foreignKey:CustomerID"`

	MenuItems []Menu `json:"menu_items" gorm:"many2many:order_menu_items;"`

	Status     string `json:"status"` // "pending", "in_progress", "done"
	Note       string `json:"note"`   // optional: e.g., "no sambal"
	TotalPrice int    `json:"total_price"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Order) TableName() string {
	return "order"
}
