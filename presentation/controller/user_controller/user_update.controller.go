package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/domain/contracts/user_contracts"
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/config"
	"github.com/junioralcant/api-stores-go/presentation/helpers"
	"github.com/junioralcant/api-stores-go/validation/user_validation"
)

type UserUpdateController struct {
	UseCase user_contracts.IUserUpdate
}

func (u *UserUpdateController) Handle(ctx *gin.Context) {

	request := user_validation.UpdateUser{}

	ctx.BindJSON(&request.User)

	if err := request.Validate(); err != nil {
		config.Log.Errorf("Error validating request: %v", err.Error())
		helpers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		helpers.SendError(ctx, http.StatusBadRequest, helpers.ErrParamIsRequired("id", "string").Error())
		return
	}

	user := models.User{
		Name:     request.User.Name,
		Email:    request.User.Email,
		Password: request.User.Password,
		Phone:    request.User.Phone,
		CPF:      request.User.CPF,
	}

	response, err := u.UseCase.Update(id, user)

	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SendSuccess(ctx, "user updated", response)
}
