package user_repository

import (
	"fmt"

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

func (r *UserRepository) Update(id string, user models.User) (*models.User, error) {

	userUpdated := models.User{}

	if err := config.DB.First(&userUpdated, id).Error; err != nil {
		return nil, fmt.Errorf("error in search user: %+v", err)
	}

	if user.Name != "" {
		userUpdated.Name = user.Name
	}

	if user.Email != "" {
		userUpdated.Email = user.Email
	}

	if user.Password != "" {
		userUpdated.Password = user.Password
	}

	if user.Phone != "" {
		userUpdated.Phone = user.Phone
	}

	if user.CPF != "" {
		userUpdated.CPF = user.CPF
	}

	if err := config.DB.Save(&userUpdated).Error; err != nil {
		return nil, fmt.Errorf("error in update user: %+v", err)
	}

	return &userUpdated, nil

}
