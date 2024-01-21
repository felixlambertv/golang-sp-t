package model

type User struct {
	Id          uint   `json:"-"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"-"`
}
