package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phonghaido/api-gateway/handlers"
	"github.com/phonghaido/api-gateway/helpers"
	l "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	l.Infoln("Artifactory upload service is listening at port 8080...")

	userGroup := e.Group("/api/users")

	userGroup.POST("/add", helpers.EchoErrorWrapper(handlers.HandlePostAddUser))

	e.Logger.Fatal(e.Start(":3000"))
}
