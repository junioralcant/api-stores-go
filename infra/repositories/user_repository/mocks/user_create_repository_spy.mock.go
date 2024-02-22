package mocks

import "github.com/junioralcant/api-stores-go/domain/models"

type UserCreateRepositorySpy struct {
	User models.User
}

func NewUserCreateRepositorySpy() *UserCreateRepositorySpy {
	return &UserCreateRepositorySpy{}
}

func (r *UserCreateRepositorySpy) UserCreateRepo(user models.User) *models.User {
	r.User = user
	return &user
}
