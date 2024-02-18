package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   bool   `json:"status" gorm:"default:true"`
}
