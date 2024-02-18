package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/contracts/user_contracts_infra"
)

type UserUpdateUseCase struct {
	Repo user_contracts_infra.IUserRepository
}

func (u *UserUpdateUseCase) Update(id string, user models.User) (*models.User, error) {
	userCreated, err := u.Repo.Update(id, user)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}
