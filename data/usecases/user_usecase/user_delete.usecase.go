package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
)

type UserDeleteUseCase struct {
	Repo *user_repository.UserRepository
}

func (uc *UserDeleteUseCase) UserDelete(id string) (*models.User, error) {

	userDelete, err := uc.Repo.Delete(id)
	if err != nil {
		return nil, err
	}

	return userDelete, nil
}
