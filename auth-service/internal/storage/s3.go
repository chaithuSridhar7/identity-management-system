package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Storage() (*S3Storage, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	bucketName := os.Getenv("AWS_S3_BUCKET")
	if bucketName == "" {
		return nil, fmt.Errorf("AWS_S3_BUCKET is not set")
	}

	return &S3Storage{
		Client:     s3.NewFromConfig(cfg),
		BucketName: bucketName,
	}, nil
}

func (s *S3Storage) Upload(
	ctx context.Context,
	file multipart.File,
	key string,
	contentType string,
) error {
	_, err := s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.BucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})

	return err
}

func (s *S3Storage) CheckBucket(ctx context.Context) error {
	_, err := s.Client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(s.BucketName),
	})

	return err
}

func (s *S3Storage) Download(
	ctx context.Context,
	key string,
) (*s3.GetObjectOutput, error) {
	return s.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(key),
	})
}
