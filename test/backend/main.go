package main

import (
	"log"
	"net/http"
	"strconv" // Add the import statement for strconv

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type jsonData struct {
	gorm.Model
	Name   int    `json:"name"`
	Descri string `json:"descri"`
	Bool   bool   `json:"bool"`
}

var db *gorm.DB

func main() {

	var err error
	dsn := "user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"}, // フロントエンドのURL
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	// マイグレーション: データベースにjsonDataテーブルを作成
	db.AutoMigrate(&jsonData{})

	var count int64
    db.Model(&jsonData{}).Count(&count)
    if count == 0 {
        // データベースが空であれば、初期データを挿入
        db.Create(&jsonData{
            Name:   1,
            Descri: "Initial data",
            Bool:   true,
        })
    }

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/add", addHandler)
	e.GET("/json", jsonHandler)
	e.DELETE("/delete/:name", deleteHandler)

	e.Logger.Fatal(e.Start(":8081"))
}

func deleteHandler(c echo.Context) error {
	name := c.Param("name")
	intName, err := strconv.Atoi(name)
	if err != nil {
		log.Printf("Invalid name: %v", mysql.DefaultDriverName)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid name")
	}

	 // データベースからレコードを検索
	 var jsonDataItem jsonData
	 result := db.First(&jsonDataItem, intName)

	 if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            log.Printf("ID not found: %v", name)
            return echo.NewHTTPError(http.StatusNotFound, "ID not found")
        }
        log.Printf("Failed to retrieve record: %v", result.Error)
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve record")
    }


	if err := db.Delete(&jsonData{}, intName).Error; err != nil {
		log.Printf("Failed to delete data with ID %v: %v", name, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete data")
	}

	log.Printf("Delete ID: %v", name)

	// Correctly adjust index as IDs are 1-based
	return c.JSON(http.StatusOK, "Data deleted successfully")
}

func jsonHandler(c echo.Context) error {
	var dataList []jsonData
	if err := db.Find(&dataList).Error; err != nil {
		log.Printf("Failed to retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve data")
	}

	log.Printf("Data List: %v", dataList)
	return c.JSON(http.StatusOK, dataList)
}
func addHandler(c echo.Context) error {
	var newdata jsonData
	if err := c.Bind(&newdata); err != nil {
		return err
	}

	// データベースに新しいレコードを追加
	if err := db.Create(&newdata).Error; err != nil {
		log.Printf("Failed to create record: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create record")
	}

	return c.JSON(http.StatusOK, newdata)
}
