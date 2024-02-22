package tests

import (
	"testing"

	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository/mocks"
)

func TestUserCreateUseCase_should_call_UserCreate_params_with_correct_values(t *testing.T) {
	repoSpy := mocks.NewUserCreateRepositorySpy()
	useCase := user_usecase.NewUserCreateUseCase(repoSpy)

	user := models.User{
		ID:       1,
		Name:     "jhon doe",
		Email:    "jhon@example.com",
		Phone:    "9999",
		CPF:      "111.111.111-11",
		Password: "123456",
	}

	useCase.UserCreate(user)

	if repoSpy.User.Name != user.Name {
		t.Errorf("expected %s, got %s", user.Name, repoSpy.User.Name)
	}

	if repoSpy.User.Email != user.Email {
		t.Errorf("expected %s, got %s", user.Email, repoSpy.User.Email)
	}

	if repoSpy.User.Phone != user.Phone {
		t.Errorf("expected %s, got %s", user.Phone, repoSpy.User.Phone)
	}

	if repoSpy.User.CPF != user.CPF {
		t.Errorf("expected %s, got %s", user.CPF, repoSpy.User.CPF)
	}

	if repoSpy.User.Password != "user.Password" {
		t.Errorf("expected %s, got %s", "user.Password", repoSpy.User.Password)
	}

	if repoSpy.User.ID != user.ID {
		t.Errorf("expected %d, got %d", user.ID, repoSpy.User.ID)
	}
}
