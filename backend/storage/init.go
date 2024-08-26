package storage

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type client struct {
	client *s3.Client
}

var c client

func Init() error {
	accessKey := os.Getenv("STORAGE_ACCESS_KEY")
	secret := os.Getenv("STORAGE_SECRET")
	endpoint := os.Getenv("STORAGE_BASE_ENDPOINT")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secret,
			"",
		)),
		config.WithRegion("auto"),
	)
	if err != nil {
		return err
	}

	c.client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})

	return nil
}
