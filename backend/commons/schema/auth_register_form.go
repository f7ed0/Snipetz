package schema

import "snipetz/commons/utils"

type AuthRegisterForm struct {
	Email string `json:"email"`
	AuthLoginForm
}

func (r AuthRegisterForm) AllFieldValid() bool {
	return r.AuthLoginForm.AllFieldValid() && utils.IsValidEmail(r.Email)
}
