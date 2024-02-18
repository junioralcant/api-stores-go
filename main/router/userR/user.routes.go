package userR

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	user_list_controller "github.com/junioralcant/api-stores-go/presentation/controller/user"
)

func InitUserRoutes(r *gin.Engine, apiPrefix string) {

	repo := user_repository.UserRepository{}
	useCase := user_usecase.UserListUseCase{Repo: &repo}
	userListController := user_list_controller.UserListController{UseCase: &useCase}

	g := r.Group(apiPrefix)

	g.GET("/users", userListController.Handle)
}
