package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	eMiddleware "github.com/oapi-codegen/echo-middleware"
	"net/http"
	"strings"
)

func getJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")
	// Check for the Authorization header.
	if authHdr == "" {
		return "", errors.New("ErrInvalidAuthHeader")
	}
	// We expect a header value of the form "Bearer <token>", with 1 space after
	// Bearer, per spec.
	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", errors.New("ErrInvalidAuthHeader")
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}

func JWTAuthenticationFunc(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	tokenString, err := getJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return err
	}
	secret := []byte("secret")
	claims := &handler.TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid auth token")
	}

	if claims, ok := token.Claims.(*handler.TokenClaims); ok && token.Valid {
		user := claims.User
		eCtx := eMiddleware.GetEchoContext(ctx)
		eCtx.Set("user", user)
		return nil
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid auth token")
	}
}
