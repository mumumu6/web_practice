package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"

	"git.trap.jp/1-Monthon_24_05/leaQ/backend/model"
	"git.trap.jp/1-Monthon_24_05/leaQ/backend/storage"
)

func CreatePosts(c echo.Context) error {
	log.Printf("Posts start")

	// フォームデータから他のデータを取得
	authorID := c.FormValue("author_id")
	if authorID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Author ID is required"})
	}

	description := c.FormValue("description")
	tags := c.FormValue("tags") // フォームデータ内の "tags" は単純な文字列として受け取ります

	// JSON文字列をスライスに変換
	var tagsSlice []string
	if err := json.Unmarshal([]byte(tags), &tagsSlice); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	var tagModels []model.TagsData
	for _, tagName := range tagsSlice {
		tag := model.TagsData{Tag: tagName}
		// 既存のタグがある場合はそれを使用し、なければ新規作成
		model.DB.Where("tag = ?", tagName).FirstOrCreate(&tag)
		tagModels = append(tagModels, tag)
		log.Printf("Tag: %v", tag)
	}
	log.Printf("Tags: %v", tagModels)

	// フォームデータからファイルを取得
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	// ファイルの種類とサイズを検証
	if file.Size > 10*1024*1024 { // 10MBを超えるファイルは拒否
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "File size exceeds limit"})
	}

	fileName, err := storage.Upload(file)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	postData := model.PostsData{
		AuthorID:    authorID,
		Tags:        tagModels,
		Description: description,
		ImageName:   fileName, // 画像パスを設定
	}

	log.Printf("Post data: %v", postData)

	if model.DB == nil {
		log.Printf("Database connection is nil")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	// データベースに保存
	if err := model.DB.Create(&postData).Error; err != nil {
		log.Printf("Failed to save data: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	log.Printf("Saved data: %v", postData)

	return c.JSON(http.StatusOK, postData)
}

func GetRecentPosts(c echo.Context) error {
	log.Printf("GetRecentPosts start")

	// データベースから最新の投稿10件を取得（タグもプレロード）
	var posts []model.PostsData
	if err := model.DB.Preload("Tags").
		Preload("Comments"). // コメントもプレロード
		Order("created_at DESC").
		Limit(10).
		Find(&posts).Error; err != nil {
		log.Printf("Failed to get data: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	// 投稿データを整形して返す
	var response []map[string]interface{}
	for _, post := range posts {
		sort.Slice(post.Comments, func(i, j int) bool {
			return post.Comments[i].CreatedAt.Before(post.Comments[j].CreatedAt)
		})
		response = append(response, map[string]interface{}{
			"post_id":     post.ID, // コメントの際投稿を識別するためのID
			"author_id":   post.AuthorID,
			"description": post.Description,
			"image_name":  post.ImageName,
			"tags":        post.Tags,     // タグ情報
			"comments":    post.Comments, // コメント情報
			"CreatedAt":   post.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, response)
}

// getImageURLは画像のパスからフルURLを生成します。
// func getImageURL(imagePath string) string {
// 	// ここで画像のベースURLを設定
// 	baseURL := "https://yourdomain.com/images/"
// 	return baseURL + imagePath
// }

func GetUserPosts(c echo.Context) error {
	log.Printf("GetUserPosts start")

	// ユーザーIDを取得
	userID := c.Param("user_id")

	// データベースからユーザーの投稿を取得（タグもプレロード）
	var posts []model.PostsData
	if err := model.DB.Preload("Tags").
		Preload("Comments"). // コメントもプレロード
		Where("author_id = ?", userID).
		Order("created_at DESC").
		Limit(10).
		Find(&posts).Error; err != nil {
		log.Printf("Failed to get data: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Internal Server Error"})
	}

	// 投稿データを整形して返す
	var response []map[string]interface{}
	for _, post := range posts {
		sort.Slice(post.Comments, func(i, j int) bool {
			return post.Comments[i].CreatedAt.Before(post.Comments[j].CreatedAt)
		})
		response = append(response, map[string]interface{}{
			"post_id":     post.ID,
			"author_id":   post.AuthorID,
			"description": post.Description,
			"image_name":  post.ImageName,
			"tags":        post.Tags, // タグ情報
			"comments":    post.Comments,
			"CreatedAt":   post.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, response)
}
