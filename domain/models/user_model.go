package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id" gorm:"type:uuid; primarykey;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    *bool     `json:"status" gorm:"default:true"`
}

func NewUser(name string, email string, phone string, cpf string, password string, status *bool) *User {
	user := User{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Phone:    phone,
		CPF:      cpf,
		Password: password,
		Status:   status,
	}

	return &user
}
