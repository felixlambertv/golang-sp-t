package handler

import (
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang-jwt/jwt"
)

type UserRegisterRequest struct {
	PhoneNumber string
	FullName    string
	Password    string
}

type TokenClaims struct {
	jwt.StandardClaims
	User   *repository.UserOutput
	Expire int64
}
