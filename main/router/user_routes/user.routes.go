package user_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/main/factories/user_factories"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func InitUserRoutes(r *gin.Engine, apiPrefix string) {

	repo := user_repository.UserRepository{}

	useCaseUpdate := user_usecase.UserUpdateUseCase{Repo: &repo}
	userUpdateController := user_controller.UserUpdateController{UseCase: &useCaseUpdate}

	g := r.Group(apiPrefix)

	g.GET("/users", user_factories.UserListAllControllerFactory().Handle)

	g.POST("/user", user_factories.UserCreateControllerFactory().Handle)

	g.PUT("/user", userUpdateController.Handle)

	g.DELETE("/user", user_factories.UserDeleteControllerFactory().Handle)
}
