package main

import (
	"os"

	"git.trap.jp/1-Monthon_24_05/leaQ/backend/handler"
	"git.trap.jp/1-Monthon_24_05/leaQ/backend/model"
	"git.trap.jp/1-Monthon_24_05/leaQ/backend/storage"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
	}

	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	err := model.Init()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 「/hello」というエンドポイントを設定する
	e.GET("/hello", handler.Helloworld)
	e.GET("/api/getrecentposts", handler.GetRecentPosts)
	e.GET("/api/:user_id/posts", handler.GetUserPosts)
	e.GET("/api/images/:image_name", handler.GetImage)
	e.GET("/api/tags/:tag_name", handler.GetTagPosts)
	e.POST("/api/posts", handler.CreatePosts)
	e.POST("/api/comments", handler.CreateComments)
	e.POST("/api/bookmark", handler.CreateBookmarks)	

	addr := os.Getenv("SERVER_ADDRESS")
	e.Logger.Fatal(e.Start(addr))
}
