package user_repository

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/config"
)

type UserRepository struct {
}

func (r *UserRepository) ListAll() []models.User {

	users := []models.User{}

	if err := config.DB.Find(&users).Error; err != nil {
		config.Log.Errorf("error in list openings: %v", err)
	}

	return users
}

func (r *UserRepository) Create(user models.User) *models.User {
	userCreate := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
		CPF:      user.CPF,
	}

	if err := config.DB.Create(&userCreate).Error; err != nil {
		config.Log.Errorf("error in create user: %v", err)
		return nil
	}

	return &userCreate
}
