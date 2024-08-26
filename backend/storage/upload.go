package storage

import (
	"context"
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func Upload(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// UUIDを使ってユニークなファイル名を生成
	extension := filepath.Ext(fileHeader.Filename) // 元のファイルの拡張子を取得
	if extension == "" {
        return "", errors.New("invalid file extension")
    } // 拡張子がない場合はエラーを返す

	// 許可された拡張子のリスト
    allowedExtensions := map[string]bool{
        ".jpg":  true,
        ".jpeg": true,
        ".png":  true,
        ".gif":  true,
    }

	if !allowedExtensions[strings.ToLower(extension)] {
        return "", errors.New("file type not allowed")
    } // 許可された拡張子以外のファイルはエラーを返す

	
	uniqueUUID, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	uniqueFileName := uniqueUUID.String() + extension
	
	if _, err := c.client.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String("leaq"),
			Key:    aws.String(uniqueFileName),
			Body:   file,
		},
	); err != nil {
		return "", err
	}

	return uniqueFileName, nil
}
