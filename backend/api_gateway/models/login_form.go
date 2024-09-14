package models

import "snipetz/commons/schema"

type LoginForm struct {
	Ip string `json:"ip"`
	schema.AuthLoginForm
}

func (l LoginForm) AllFieldValid() bool {
	return l.AuthLoginForm.AllFieldValid() && l.Ip != ""
}
