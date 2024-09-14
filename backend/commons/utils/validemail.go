package utils

import "regexp"

func IsValidEmail(email string) (ok bool) {
	ok, err := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", email)
	if err != nil {
		return false
	}
	return
}
