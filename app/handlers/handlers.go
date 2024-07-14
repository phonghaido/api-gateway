package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phonghaido/api-gateway/helpers"
)

func HandlePostAddUser(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return helpers.MethodNotAllowed()
	}
	return nil
}
