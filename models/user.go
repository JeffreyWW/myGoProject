package models

type User struct {
	Id       int    `json:"userId"`
	PhoneNo  int    `json:"phoneNo"`
	Password string `json:"-"`
}
