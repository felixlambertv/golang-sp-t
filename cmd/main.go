package main

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	middleware2 "github.com/SawitProRecruitment/UserService/middleware"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	eMiddleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	var server generated.ServerInterface = newServer()

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("api1.yaml")
	if err != nil {
		log.Fatalf("Failed to load OpenAPI spec: %v", err)
	}
	e.Use(eMiddleware.OapiRequestValidatorWithOptions(doc, &eMiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: middleware2.JWTAuthenticationFunc,
		},
		SilenceServersWarning: true,
	}))

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
