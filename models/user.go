package models

type User struct {
	Id       uint   `json: "id"`
	Email    string `json: "email"`
	Password []byte `json: "-"`
}