package user_list_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/domain/contracts/user_contracts"
)

type UserListController struct {
	UseCase user_contracts.IUserList
}

func (u *UserListController) Handle(ctx *gin.Context) {
	users := u.UseCase.FindAll()
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"data": users,
	})
}
