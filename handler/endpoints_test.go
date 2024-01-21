package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("Successful Login", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := repository.NewMockRepositoryInterface(ctrl)
		e := echo.New()

		password := "password"
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		assert.NoError(t, err)

		userOutput := &repository.UserOutput{
			Id:          1,
			PhoneNumber: "123456789",
			Password:    string(hash),
		}

		loginReq := generated.LoginRequest{
			PhoneNumber: "123456789",
			Password:    password,
		}
		requestBody, _ := json.Marshal(loginReq)
		req := httptest.NewRequest(echo.POST, "/login", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		s := &Server{Repository: mockRepo}

		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(userOutput, nil)
		mockRepo.EXPECT().IncrementUserLoginCount(gomock.Any(), gomock.Any()).Times(1)

		assert.NoError(t, s.Login(c))
		var loginResponse generated.LoginResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &loginResponse))

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, int(userOutput.Id), loginResponse.UserId)
		assert.NotEmpty(t, loginResponse.Token)
	})

	t.Run("Incorrect Password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := repository.NewMockRepositoryInterface(ctrl)
		e := echo.New()

		password := "password"
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		assert.NoError(t, err)

		userOutput := &repository.UserOutput{
			Id:          1,
			PhoneNumber: "123456789",
			Password:    string(hash),
		}

		loginReq := generated.LoginRequest{
			PhoneNumber: "123456789",
			Password:    "wrongpass",
		}
		requestBody, _ := json.Marshal(loginReq)
		req := httptest.NewRequest(echo.POST, "/login", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		s := &Server{Repository: mockRepo}

		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(userOutput, nil)

		assert.NoError(t, s.Login(c))
		var resp generated.GeneralErrorResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "Invalid credential", resp.Message)
	})

	t.Run("Error Incrementing User Login Count", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := repository.NewMockRepositoryInterface(ctrl)
		e := echo.New()

		password := "password"
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		assert.NoError(t, err)

		userOutput := &repository.UserOutput{
			Id:          1,
			PhoneNumber: "123456789",
			Password:    string(hash),
		}

		loginReq := generated.LoginRequest{
			PhoneNumber: "123456789",
			Password:    password,
		}
		requestBody, _ := json.Marshal(loginReq)
		req := httptest.NewRequest(echo.POST, "/login", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		s := &Server{Repository: mockRepo}

		errRepo := errors.New("error increment")
		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(userOutput, nil)
		mockRepo.EXPECT().IncrementUserLoginCount(gomock.Any(), gomock.Any()).Return(errRepo)

		assert.NoError(t, s.Login(c))
		var resp generated.GeneralErrorResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Something went wrong", resp.Message)
	})
}

func TestGetProfile(t *testing.T) {
	e := echo.New()
	userOutput := &repository.UserOutput{}

	t.Run("User Found", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/profile", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user", userOutput)

		s := &Server{}
		assert.NoError(t, s.GetProfile(c))
		var resp repository.UserOutput
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

		assert.Equal(t, userOutput.Id, resp.Id)
		assert.Equal(t, userOutput.PhoneNumber, resp.PhoneNumber)
	})

	t.Run("User Not Found", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/profile", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user", nil)

		s := &Server{}
		assert.NoError(t, s.GetProfile(c))
		var resp generated.GeneralErrorResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, "Unauthorized", resp.Message)
	})
}

func TestUserRegister(t *testing.T) {
	e := echo.New()
	t.Run("Successful Registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := repository.NewMockRepositoryInterface(ctrl)

		reqBody := generated.RegistrationRequest{
			FullName:    "John Doe",
			PhoneNumber: "123456789",
			Password:    "password",
		}
		requestBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(echo.POST, "/register", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().IsPhoneNumberExists(gomock.Any(), reqBody.PhoneNumber).Return(false, nil)
		mockUserOutput := &repository.CreateUserOutput{Id: 1}
		mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(mockUserOutput, nil)

		s := &Server{Repository: mockRepo}

		assert.NoError(t, s.UserRegister(c))
		var resp generated.RegistrationResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

		assert.Equal(t, int(mockUserOutput.Id), resp.UserId)
	})

	t.Run("Phone Number Already Exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := repository.NewMockRepositoryInterface(ctrl)

		reqBody := generated.RegistrationRequest{
			FullName:    "John Doe",
			PhoneNumber: "123456789",
			Password:    "password",
		}
		requestBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(echo.POST, "/register", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().IsPhoneNumberExists(gomock.Any(), reqBody.PhoneNumber).Return(true, nil)

		s := &Server{Repository: mockRepo}

		assert.NoError(t, s.UserRegister(c))
		var resp generated.GeneralErrorResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.Equal(t, "Phone number already exist", resp.Message)
	})

	t.Run("Error on Phone Number Existence Check", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := repository.NewMockRepositoryInterface(ctrl)

		reqBody := generated.RegistrationRequest{
			FullName:    "John Doe",
			PhoneNumber: "123456789",
			Password:    "password",
		}
		requestBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(echo.POST, "/register", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().IsPhoneNumberExists(gomock.Any(), reqBody.PhoneNumber).Return(false, errors.New("database error"))

		s := &Server{Repository: mockRepo}

		assert.NoError(t, s.UserRegister(c))
		var resp generated.GeneralErrorResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	})

	t.Run("Error on User Creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := repository.NewMockRepositoryInterface(ctrl)

		reqBody := generated.RegistrationRequest{
			FullName:    "John Doe",
			PhoneNumber: "123456789",
			Password:    "password",
		}
		requestBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(echo.POST, "/register", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().IsPhoneNumberExists(gomock.Any(), reqBody.PhoneNumber).Return(false, nil)
		mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("creation error"))

		s := &Server{Repository: mockRepo}

		assert.NoError(t, s.UserRegister(c))
		var resp generated.GeneralErrorResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	})
}
