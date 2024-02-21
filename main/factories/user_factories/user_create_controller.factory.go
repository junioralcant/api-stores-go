package user_factories

import (
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func UserDeleteControllerFactory() *user_controller.UserDeleteController {
	repo := user_repository.UserDeleteRepository{}
	useCaseCreate := user_usecase.UserDeleteUseCase{Repo: &repo}
	userDeleteController := user_controller.UserDeleteController{UseCase: &useCaseCreate}

	return &userDeleteController
}
