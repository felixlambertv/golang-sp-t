// This file contains types that are used in the repository layer.
package repository

type CreateUserInput struct {
	FullName    string
	PhoneNumber string
	Password    string
}

type UserOutput struct {
	Id          uint   `json:"-"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"-"`
}

type CreateUserOutput struct {
	Id uint
}

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}
