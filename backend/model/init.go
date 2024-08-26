package model

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PostsData struct {
	gorm.Model
	AuthorID    string          `json:"author_id"`
	Tags        []TagsData      `json:"tags" gorm:"many2many:post_tags;"` // many2manyリレーションを定義
	Description string          `json:"description"`
	ImageName   string          `json:"image_name"` // 画像のファイル名を記録
	Comments    []CommentsData  `json:"comments" gorm:"foreignKey:PostID;references:ID"`
	Bookmarks   []BookmarksData `json:"bookmarks" gorm:"foreignKey:PostID;references:ID"`
}

type BookmarksData struct {
	gorm.Model
	PostID uint   `json:"post_id"`
	UserID string `json:"user_id"` 
}

type TagsData struct {
	gorm.Model
	Tag string `json:"tag"`
}

type CommentsData struct {
	gorm.Model
	PostID   uint   `json:"post_id"`
	AuthorID string `json:"author_id"`
	Comment  string `json:"comment"`
}

var user = "root"
var pass = "password"
var host = "localhost"
var dbname = "database"

var DB *gorm.DB

func Init() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, dbname) + "?parseTime=True&loc=Asia%2FTokyo&charset=utf8mb4"
	tempDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	// モデルをマイグレーション
	if err := tempDB.AutoMigrate(&PostsData{}, &TagsData{}, &CommentsData{}, &BookmarksData{}); err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	DB = tempDB
	log.Printf("Connected to database: %s", dbname)
	return nil
}
