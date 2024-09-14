package schema

type AuthLoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l AuthLoginForm) AllFieldValid() bool {
	return l.Username != "" && l.Password != ""
}
