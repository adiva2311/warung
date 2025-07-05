package services

import (
	"errors"
	"log"
	"warung/helpers"
	"warung/models"
	"warung/repositories"

	"gorm.io/gorm"
)

type UserService interface {
	Register(request models.User) (helpers.UserResponse, error)
	Login(request helpers.LoginRequest) (helpers.LoginResponse, error)
}

type UserServiceImpl struct {
	UserRepo repositories.UserRepository
}

// Login implements UserService.
func (u *UserServiceImpl) Login(request helpers.LoginRequest) (helpers.LoginResponse, error) {
	// Check Email if Exist
	user, err := u.UserRepo.CheckEmail(request.Email)
	if err != nil {
		return helpers.LoginResponse{}, errors.New("invalid username or password")
	}

	// Check Password
	if !helpers.CheckPasswordHash(request.Password, user.Password) {
		return helpers.LoginResponse{}, errors.New("invalid username or password")
	}

	// Generate JWT
	token, err := helpers.GenerateJWT(int(user.ID), user.Email, user.Role, user.PhoneNumber)
	if err != nil {
		return helpers.LoginResponse{}, err
	}
	return helpers.ToLoginResponse(user, token), nil
}

// Register implements UserService.
func (u *UserServiceImpl) Register(request models.User) (helpers.UserResponse, error) {
	hashPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		log.Fatal(err)
	}

	user := &models.User{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Role:        request.Role,
		Password:    hashPassword,
	}

	err = u.UserRepo.Register(*user)
	if err != nil {
		log.Fatal("Failed to Register User")
	}

	return helpers.ToRegisterResponse(*user), nil
}

func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepo: repositories.NewUserRepository(db),
	}
}
