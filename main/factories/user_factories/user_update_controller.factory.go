package user_factories

import (
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func UserUpdateControllerFactory() *user_controller.UserUpdateController {
	repo := user_repository.UserUpdateRepository{}
	useCase := user_usecase.UserUpdateUseCase{Repo: &repo}
	userUpdateController := user_controller.UserUpdateController{UseCase: &useCase}

	return &userUpdateController
}
