package main

import (
	"log"
	"net/http"
	"strconv" // Add the import statement for strconv

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jsonData struct {
	Name   int    `json:"name"`
	Descri string `json:"descri"`
	Bool   bool   `json:"bool"`
}

var dataList []jsonData

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"}, // フロントエンドのURL
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	dataList = append(dataList, jsonData{
		Name:   1,
		Descri: "test",
		Bool:   true,
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/add", addHandler)
	e.GET("/json", jsonHandler)
	e.DELETE("/delete/:id", deleteHandler)

	e.Logger.Fatal(e.Start(":8081"))
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Invalid ID: %v", id)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if intID < 1 || intID > len(dataList) {
		log.Printf("ID not found: %v", id)
		return echo.NewHTTPError(http.StatusNotFound, "ID not found")
	}

	log.Printf("Delete ID: %v", id)

	// Correctly adjust index as IDs are 1-based
	dataList = append(dataList[:intID-1], dataList[intID:]...)
	return c.JSON(http.StatusOK, dataList)
}

func jsonHandler(c echo.Context) error {

	log.Printf("Data List: %v", dataList)

	return c.JSON(http.StatusOK, dataList)
}
func addHandler(c echo.Context) error {
	var newdata jsonData
	if err := c.Bind(&newdata); err != nil {
		return err
	}

	// 受け取ったデータを処理する
	dataList = append(dataList, newdata)
	// ここでデータベースに保存するなどの処理を行います

	return c.JSON(http.StatusOK, newdata)
}
