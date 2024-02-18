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

func (r *UserRepository) Update(id string, user models.User) *models.User {

	userUpdated := user

	if err := config.DB.First(&user, id).Error; err != nil {
		config.Log.Errorf("error in search user: %+v", err)
		return nil
	}

	userUpdated.Name = user.Name
	userUpdated.Email = user.Email
	userUpdated.Password = user.Password
	userUpdated.Phone = user.Phone
	userUpdated.CPF = user.CPF

	if err := config.DB.Save(&userUpdated).Error; err != nil {
		config.Log.Errorf("error in update user: %+v", err)
		return nil
	}

	return &userUpdated

}
