package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/contracts/user_contracts_infra"
)

type UserUpdateUseCase struct {
	Repo user_contracts_infra.IUserRepository
}

func (u *UserUpdateUseCase) Update(id string, user models.User) *models.User {
	userCreated := u.Repo.Update(id, user)
	return userCreated
}
