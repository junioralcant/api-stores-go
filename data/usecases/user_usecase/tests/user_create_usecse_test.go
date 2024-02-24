package tests

import (
	"testing"

	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository/mocks"
	"github.com/stretchr/testify/assert"
)

func makeSut() (*user_usecase.UserCreateUseCase, *mocks.UserCreateRepositorySpy) {
	repo := mocks.NewUserCreateRepositorySpy()
	useCase := user_usecase.NewUserCreateUseCase(repo)
	return useCase, repo
}

var userMocked = models.User{
	ID:       1,
	Name:     "jhon doe",
	Email:    "jhon@example.com",
	Phone:    "9999",
	CPF:      "111.111.111-11",
	Password: "123456",
}

func TestUserCreateUseCase_correct_params(t *testing.T) {
	useCase, repoSpy := makeSut()

	useCase.UserCreate(userMocked)

	if repoSpy.Params != userMocked {
		t.Errorf("expected user %v, got %v", userMocked, repoSpy.Params)
	}
}

func TestUserCreateUseCase_correct_response(t *testing.T) {
	useCase, _ := makeSut()

	userCreated := useCase.UserCreate(userMocked)

	assert.EqualValues(t, userCreated.Name, userMocked.Name)
	assert.EqualValues(t, userCreated.Email, userMocked.Email)
	assert.EqualValues(t, userCreated.Phone, userMocked.Phone)
	assert.EqualValues(t, userCreated.CPF, userMocked.CPF)
	assert.EqualValues(t, userCreated.Password, userMocked.Password)
	assert.EqualValues(t, userCreated.ID, userMocked.ID)
}
