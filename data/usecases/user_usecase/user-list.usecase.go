package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/contracts/user_contracts_infra"
)

type UserListUseCase struct {
	Repo user_contracts_infra.IUserRepository
}

func (u *UserListUseCase) FindAll() []models.User {
	users := u.Repo.ListAll()

	return users
}
