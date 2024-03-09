package user_repository

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/config"
)

type UserCreateRepository struct {
}

func NewUserCreateRepository() *UserCreateRepository {
	return &UserCreateRepository{}
}

func (r *UserCreateRepository) UserCreateRepo(user models.User) *models.User {
	userCreate := models.User{
		ID:       user.ID,
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
