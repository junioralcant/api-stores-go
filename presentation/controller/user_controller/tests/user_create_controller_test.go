package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase/mocks"
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestUserCreateController_Handle(t *testing.T) {
	usecase := mocks.NewUserCreateUseCaseSpy()
	controller := user_controller.NewUserCreateController(usecase)

	recorder := httptest.NewRecorder()

	context := GetTestGinContext(recorder)

	url := url.Values{}

	body := models.User{
		ID:       "19990",
		Name:     "test",
		Email:    "junior@test",
		Password: "12345",
		Phone:    "9999999999",
		CPF:      "48893993",
	}

	MakeHttpRequest(context, url, body, gin.Params{})

	controller.Handle(context)

	response := recorder.Body.String()
	data := gjson.Get(response, "data")

	assert.EqualValues(t, http.StatusOK, recorder.Code)
	assert.EqualValues(t, data.Get("name").String(), body.Name)
	assert.EqualValues(t, data.Get("email").String(), body.Email)
	assert.EqualValues(t, data.Get("password").String(), body.Password)
	assert.EqualValues(t, data.Get("phone").String(), body.Phone)
	assert.EqualValues(t, data.Get("cpf").String(), body.CPF)
}
