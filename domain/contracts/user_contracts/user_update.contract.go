package user_contracts

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserUpdate interface {
	Update(id string, user models.User) *models.User
}
