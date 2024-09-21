package utils

import (
	"regexp"

	"github.com/f7ed0/golog/lg"
)

func IsUsernameValid(username string) bool {
	ok, err := regexp.Match(`(([_]+)?([a-z0-9]+)[\.]?)+([_]+)?$`, []byte(username))
	if err != nil {
		lg.Error.Println("Error during regex :", err.Error())
		return false
	}
	return ok && len(username) >= 5 && len(username) <= 20
}
