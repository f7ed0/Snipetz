package schema

import "snipetz/commons/utils"

// ------------------------------------------------

type AuthLoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l AuthLoginForm) AllFieldValid() bool {
	return utils.IsUsernameValid(l.Username) && l.Password != ""
}

// -------------------------------------------------

type AuthRegisterForm struct {
	Email string `json:"email"`
	AuthLoginForm
}

func (r AuthRegisterForm) AllFieldValid() bool {
	return r.AuthLoginForm.AllFieldValid() && utils.IsValidEmail(r.Email)
}

// --------------------------------------------------

type AuthRegisterResponse struct {
	Status        string `json:"status"` // valid, error
	Uuid          string `json:"uuid,omitempty"`
	InvalidReason string `json:"invalid_reason,omitempty"`
	DefaultMSResponse
}

// ---------------------------------------------------

type AuthLoginResponse struct {
	Status        string `json:"status"`
	Uuid          string `json:"uuid,omitempty"`
	InvalidReason string `json:"invalid_reason,omitempty"`
	JwtToken      string `json:"jwt_token,omitempty"`
}

// -----------------------------------------------------
