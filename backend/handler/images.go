package handler

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"git.trap.jp/1-Monthon_24_05/leaQ/backend/storage"
)

func GetImage(c echo.Context) error {
	imageName := c.Param("image_name")

	obj, err := storage.Download(imageName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	resp := c.Response()
	io.Copy(resp.Writer, obj.Body)

	return c.NoContent(http.StatusOK)
}
