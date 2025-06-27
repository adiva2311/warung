package models

type Category struct {
	ID           uint   `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	CategoryName string `json:"category_name" gorm:"column:category_name"`
	MenuItems    []Menu `gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string {
	return "category"
}
