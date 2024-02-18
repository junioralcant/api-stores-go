package user_contracts

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserList interface {
	FindAll() []models.User
}
