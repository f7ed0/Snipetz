package snipetzjwt

import (
	"os"
	snipetzerror "snipetz/commons/errors"
)

var jwtSecret string = ""

func LoadSecret() error {
	tok := os.Getenv("jwt_secret")
	if tok == "" {
		return snipetzerror.ErrorNoJwtSecret
	}
	jwtSecret = tok

	return nil

}

func GetSecret() string {
	return jwtSecret
}
