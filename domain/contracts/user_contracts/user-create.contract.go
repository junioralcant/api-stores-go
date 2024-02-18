package user_contracts

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserCreate interface {
	Create(user models.User) *models.User
}
