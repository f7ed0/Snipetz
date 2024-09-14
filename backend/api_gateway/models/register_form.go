package models

import (
	"snipetz/commons/schema"
	"snipetz/commons/utils"
)

type RegisterForm struct {
	Email string `json:"email"`
	schema.AuthRegisterForm
}

func (r RegisterForm) AllFieldValid() bool {
	return r.AuthLoginForm.AllFieldValid() && utils.IsValidEmail(r.Email)
}
