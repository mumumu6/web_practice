package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Helloworld(c echo.Context) error {
	log.Printf("Hello, World!")
	return c.String(http.StatusOK, "Hello, World!")
}