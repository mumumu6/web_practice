package handler

import (
	"log"
	"net/http"
	"strconv"

	"git.trap.jp/1-Monthon_24_05/leaQ/backend/model" // Replace "your-package" with the actual package name
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


func CreateBookmarks(c echo.Context) error {
	log.Printf("Bookmarks start")

	// フォームデータから他のデータを取得
	userID := c.FormValue("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "User ID is required"})
	}

	postIDStr := c.FormValue("post_id")
	if postIDStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Post ID is required"})
	}
	
	postID ,err :=strconv.ParseUint(postIDStr, 10, 32) 

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Post ID is invalid"})
	}

	// 同じ user_id と post_id の組み合わせが存在するか確認
    if err := model.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&model.BookmarksData{}).Error; err == nil {
        // 既に存在する場合
		log.Printf("Bookmark already exists")
        return c.JSON(http.StatusConflict, echo.Map{"error": "Bookmark already exists"})
    } else if  err != gorm.ErrRecordNotFound {
        // データベースエラーが発生した場合
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
    }

    // 存在しない場合、新しいデータを作成
    bookmarkData := model.BookmarksData{
        UserID: userID,
        PostID: uint(postID),
    }

    if err := model.DB.Create(&bookmarkData).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create bookmark"})
    }

	
	

	
	return c.JSON(http.StatusOK, echo.Map{"message": "Bookmark created"})
}