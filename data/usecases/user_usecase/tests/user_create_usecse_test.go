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

func TestUserCreateUseCase_should_call_UserCreate_params_with_correct_values(t *testing.T) {
	useCase, repoSpy := makeSut()

	useCase.UserCreate(userMocked)

	if repoSpy.User != userMocked {
		t.Errorf("expected user %v, got %v", userMocked, repoSpy.User)
	}
}

func TestUserCreateUseCase_should_return_user_if_create_user_ok(t *testing.T) {
	useCase, _ := makeSut()

	userCreated := useCase.UserCreate(userMocked)

	assert.EqualValues(t, userCreated.Name, userMocked.Name)

	assert.EqualValues(t, userCreated.Email, userMocked.Email)

	assert.EqualValues(t, userCreated.Phone, userMocked.Phone)

	assert.EqualValues(t, userCreated.CPF, userMocked.CPF)

	assert.EqualValues(t, userCreated.Password, userMocked.Password)

	assert.EqualValues(t, userCreated.ID, userMocked.ID)
}
