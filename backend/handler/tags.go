package handler

import (
	"net/http"

	"git.trap.jp/1-Monthon_24_05/leaQ/backend/model"
	"github.com/labstack/echo/v4"
)

func GetTagPosts(c echo.Context) error {

	tagName := c.Param("tag_name")

	var posts []model.PostsData
	if err := model.DB.Preload("Tags").
		Preload("Comments").
		Joins("JOIN post_tags ON post_tags.posts_data_id = posts_data.id").
		Joins("JOIN tags_data ON post_tags.tags_data_id = tags_data.id").
		Where("tags_data.tag = ?", tagName).
		Order("created_at desc").
		Limit(10).
		Find(&posts).Error; err != nil {
		return c.JSON(500, echo.Map{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, posts)
}
