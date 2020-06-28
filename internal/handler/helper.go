package handler

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type APIErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func APIResponseError(c echo.Context, status int, message string) error {
	return c.JSON(status, APIErrorResponse{Code: fmt.Sprintf("%d", status), Message: message})
}

func APIResponseErrorLog(c echo.Context, status int, message string, err error) error {
	log.Println(err)
	return c.JSON(status, APIErrorResponse{Code: fmt.Sprintf("%d", status), Message: message})
}
