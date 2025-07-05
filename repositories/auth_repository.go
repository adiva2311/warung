package repositories

import (
	"warung/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CheckEmail(email string) (*models.User, error)
	Register(request models.User) error
}

type dbUser struct {
	Conn *gorm.DB
}

// CheckEmail implements UserRepository.
func (db *dbUser) CheckEmail(email string) (*models.User, error) {
	var user models.User
	err := db.Conn.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Register implements UserRepository.
func (db *dbUser) Register(request models.User) error {
	return db.Conn.Create(&request).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &dbUser{Conn: db}
}
