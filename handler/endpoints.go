package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (s *Server) Login(ctx echo.Context) error {
	var req generated.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.GeneralErrorResponse{Message: "Bad request"})
	}

	user, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), req.PhoneNumber)
	if err != nil || user == nil {
		return ctx.JSON(http.StatusConflict, generated.GeneralErrorResponse{Message: "Duplicate phone number"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.GeneralErrorResponse{Message: "Invalid credential"})
	}

	secretKey := []byte("secret")
	claims := TokenClaims{}
	claims.User = user
	claims.Expire = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.GeneralErrorResponse{Message: "Something went wrong"})
	}

	err = s.Repository.IncrementUserLoginCount(ctx.Request().Context(), user.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.GeneralErrorResponse{Message: "Something went wrong"})
	}

	return ctx.JSON(http.StatusCreated, generated.LoginResponse{
		Token:  tokenString,
		UserId: int(user.Id),
	})
}

func (s *Server) GetProfile(ctx echo.Context) error {
	user, ok := ctx.Get("user").(*repository.UserOutput)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, generated.GeneralErrorResponse{Message: "Unauthorized"})
	}
	return ctx.JSON(http.StatusOK, user)
}

func (s *Server) UserRegister(ctx echo.Context) error {
	var req generated.RegistrationRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if exists, err := s.Repository.IsPhoneNumberExists(ctx.Request().Context(), req.PhoneNumber); err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.GeneralErrorResponse{Message: "Something went wrong"})
	} else if exists {
		return ctx.JSON(http.StatusConflict, generated.GeneralErrorResponse{Message: "Phone number already exist"})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user, err := s.Repository.CreateUser(ctx.Request().Context(), repository.CreateUserInput{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Password:    string(hashPassword),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.GeneralErrorResponse{Message: "Something went wrong"})
	}

	return ctx.JSON(http.StatusCreated, generated.RegistrationResponse{UserId: int(user.Id)})
}
