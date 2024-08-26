package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Download(fileName string) (*s3.GetObjectOutput, error) {
	result, err := c.client.GetObject(
		context.TODO(),
		&s3.GetObjectInput{
			Bucket: aws.String("leaq"),
			Key:    aws.String(fileName),
		},
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}
