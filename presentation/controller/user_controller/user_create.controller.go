package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/domain/contracts/user_contracts"
	"github.com/junioralcant/api-stores-go/domain/models"
)

type UserCreateController struct {
	UseCase user_contracts.IUserCreate
}

func (u *UserCreateController) Handle(ctx *gin.Context) {
	request := models.User{}

	ctx.BindJSON(&request)

	user := u.UseCase.UserCreate(request)

	ctx.Header("Content-Type", "application/json")

	ctx.JSON(200, gin.H{
		"data": user,
	})
}
