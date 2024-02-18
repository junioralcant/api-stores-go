package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/domain/contracts/user_contracts"
	"github.com/junioralcant/api-stores-go/domain/models"
)

type UserUpdateController struct {
	UseCase user_contracts.IUserUpdate
}

func (u *UserUpdateController) Handle(ctx *gin.Context) {
	id := ctx.Query("id")
	request := models.User{}

	ctx.BindJSON(&request)

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
		CPF:      request.CPF,
	}

	response := u.UseCase.Update(id, user)

	ctx.Header("Content-Type", "application/json")

	ctx.JSON(200, gin.H{
		"data": response,
	})
}
