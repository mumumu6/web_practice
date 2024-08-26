package handler

import (
	"log"
	"net/http"
	"strconv"

	"git.trap.jp/1-Monthon_24_05/leaQ/backend/model"
	"github.com/labstack/echo/v4"
)

func CreateComments(c echo.Context) error {
	log.Printf("Comments start")
	postIDStr := c.FormValue("post_id")
	comment := c.FormValue("comment")
	author_id := c.FormValue("author_id")

	// post_idをuintに変換
	postIDUint64, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		return echo.NewHTTPError(400, "Invalid post_id format")
	}
	postID := uint(postIDUint64) // uint型にキャスト

	commentData := model.CommentsData{
		PostID:   postID,
		AuthorID: author_id,
		Comment:  comment,
	}

	if err := model.DB.Create(&commentData).Error; err != nil {
		log.Printf("Failed to create comment: %v", err)
		return echo.NewHTTPError(500, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, commentData)
}
