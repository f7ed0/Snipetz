package utils

import (
	"regexp"

	"github.com/f7ed0/golog/lg"
)

func IsValidEmail(email string) (ok bool) {
	ok, err := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", email)
	if err != nil {
		lg.Error.Println("Error during regex :", err.Error())
		return false
	}
	lg.Debug.Println(ok)
	return
}
