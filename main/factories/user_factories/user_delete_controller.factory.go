package user_factories

import (
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func UserCreateControllerFactory() *user_controller.UserCreateController {
	repo := user_repository.UserCreateRepository{}
	useCaseCreate := user_usecase.UserCreateUseCase{Repo: &repo}
	userCreateController := user_controller.UserCreateController{UseCase: &useCaseCreate}

	return &userCreateController
}
