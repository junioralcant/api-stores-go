package user_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func InitUserRoutes(r *gin.Engine, apiPrefix string) {

	repo := user_repository.UserRepository{}

	useCaseList := user_usecase.UserListUseCase{Repo: &repo}
	userListController := user_controller.UserListController{UseCase: &useCaseList}

	useCaseCreate := user_usecase.UserCreateUseCase{Repo: &repo}
	userCreateController := user_controller.UserCreateController{UseCase: &useCaseCreate}

	useCaseUpdate := user_usecase.UserUpdateUseCase{Repo: &repo}
	userUpdateController := user_controller.UserUpdateController{UseCase: &useCaseUpdate}

	g := r.Group(apiPrefix)

	g.GET("/users", userListController.Handle)

	g.POST("/user", userCreateController.Handle)

	g.PUT("/user", userUpdateController.Handle)
}
