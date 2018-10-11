package models

type User struct {
	Id       int    `json:"userId"`
	PhoneNo  int    `json:"phoneNo" valid:"Mobile"`
	Password string `json:"-"`
}
