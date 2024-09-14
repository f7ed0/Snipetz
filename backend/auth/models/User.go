package models

type User struct {
	Uid      string
	Username string
	Email    string
	Hash     string
	Salt     string
}
