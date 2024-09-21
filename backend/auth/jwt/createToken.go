package snipetzjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenFor(uuid string) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"uuid":   uuid,
		"expiry": time.Now().Add(12 * time.Hour).Unix(),
	})
	return tok.SignedString(GetSecret())
}
