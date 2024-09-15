package models

import (
	"snipetz/commons/crypt"
	"snipetz/commons/schema"

	"github.com/google/uuid"
)

type User struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Hash     string `json:"hash"`
	Salt     string `json:"salt"`
}

func GenerateUserFromRegisterForm(form schema.AuthRegisterForm) (u User, err error) {
	hash, salt, err := crypt.GeneratePair(form.Password)
	if err != nil {
		return User{}, err
	}
	u = User{
		Uid:      uuid.NewString(),
		Username: form.Username,
		Email:    form.Email,
		Hash:     hash,
		Salt:     salt,
	}
	return
}
