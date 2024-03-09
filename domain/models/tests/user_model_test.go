package tests

import (
	"testing"

	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/stretchr/testify/assert"
)

var userMocked = models.User{
	Name:     "jhon doe",
	Email:    "jhon@example.com",
	Phone:    "9999",
	CPF:      "111.111.111-11",
	Password: "123456",
}

func TestUserModel_NewUser(t *testing.T) {
	status := true
	user := models.NewUser(userMocked.Name, userMocked.Email, userMocked.Phone, userMocked.CPF, userMocked.Password, &status)

	assert.EqualValues(t, user.Name, userMocked.Name)
	assert.EqualValues(t, user.Email, userMocked.Email)
	assert.EqualValues(t, user.Phone, userMocked.Phone)
	assert.EqualValues(t, user.CPF, userMocked.CPF)
	assert.EqualValues(t, user.Password, userMocked.Password)
	assert.EqualValues(t, user.Status, &status)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}
