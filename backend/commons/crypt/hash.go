package crypt

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func HashPassword(password []byte, salt []byte) []byte {
	return pbkdf2.Key(password, salt, 210000, 128, sha256.New)
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func StringToByte(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func ByteToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GeneratePair(password string) (hash, salt string, err error) {
	b_salt, err := GenerateSalt()
	if err != nil {
		return
	}
	b_password := []byte(password)
	if err != nil {
		return
	}
	b_hash := HashPassword(b_password, b_salt)
	hash = ByteToString(b_hash)
	salt = ByteToString(b_salt)
	return
}
