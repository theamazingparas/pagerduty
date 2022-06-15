package utils

import (
	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, statusCode int, err error) error {
	return c.JSON(statusCode, map[string]interface{}{
		"message": err,
	})
}
