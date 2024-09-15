package models

import (
	"snipetz/commons/schema"
)

type RegisterForm struct {
	schema.AuthRegisterForm
}

func (r RegisterForm) AllFieldValid() bool {
	return r.AuthLoginForm.AllFieldValid()
}
