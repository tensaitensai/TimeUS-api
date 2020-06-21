package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type APIErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func APIResponseError(c echo.Context, status int, message string) error {
	return c.JSON(status, APIErrorResponse{Code: fmt.Sprintf("%d", status), Message: message})
}
