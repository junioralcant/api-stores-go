package user_factories

import (
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func UserListAllControllerFactory() *user_controller.UserListController {
	userListAllRepo := user_repository.UserListAllRepository{}
	useCaseList := user_usecase.UserListUseCase{Repo: &userListAllRepo}
	userListController := user_controller.UserListController{UseCase: &useCaseList}

	return &userListController
}
