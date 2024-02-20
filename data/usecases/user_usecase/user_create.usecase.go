package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/contracts/user_contracts_infra"
)

type UserCreateUseCase struct {
	Repo user_contracts_infra.IUserRepository
}

func (u *UserCreateUseCase) UserCreate(user models.User) *models.User {
	userCreated := u.Repo.Create(user)
	return userCreated
}
