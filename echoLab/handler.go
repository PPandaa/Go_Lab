package echoLab

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {

	return c.String(http.StatusOK, "This is Echo Lab!")

}
