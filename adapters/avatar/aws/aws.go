package aws

import (
	"fmt"
	"github.com/MamushevArup/ds-voice/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"os"
	"time"
)

type awsService struct {
	service    *s3.S3
	bucketName string
}

const (
	filePreSignDuration = time.Hour * 48
)

func New(cfg *config.Config) (CloudService, error) {

	// create a new session of aws
	newSession, err := session.NewSession(&aws.Config{
		Region:           &cfg.Aws.Region,
		Credentials:      credentials.NewStaticCredentials(cfg.Aws.AccessKeyID, cfg.Aws.SecretAccessKey, ""),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create session for aws s3 error: %w", err)
	}

	// create new aws service using the aws session
	service := s3.New(newSession)

	return &awsService{
		service:    service,
		bucketName: cfg.Aws.BucketName,
	}, nil
}

func (c *awsService) UploadOne(file *os.File) (string, error) {

	awsFileID := uuid.New().String()

	_, err := c.service.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(awsFileID),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file to aws bucket \nerror: %w", err)
	}

	return awsFileID, nil
}

func (c *awsService) GetOneUrl(uploadID string) (url string, err error) {

	req, _ := c.service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(uploadID),
	})

	url, err = req.Presign(filePreSignDuration)

	return
}
