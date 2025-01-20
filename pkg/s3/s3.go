package s3

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/internal/infra/env"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/log"
)

const (
	uploadPartSize = 5 * 1024 * 1024 // 5MB (S3 minimum part size)
)

type S3Interface interface {
	Upload(file *multipart.FileHeader) (string, error)
}

type S3Struct struct {
	uploader *s3manager.Uploader
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			WriteBufferSize:     64 * 1024,
			ReadBufferSize:      64 * 1024,
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 200,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
		Timeout: 45 * time.Second,
	}
}

func NewS3Client() (*S3Struct, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(env.AppEnv.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			env.AppEnv.AWSAccessKeyID,
			env.AppEnv.AWSSecretAccessKey,
			"",
		),
		HTTPClient: newHTTPClient(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.Concurrency = 2
		u.PartSize = uploadPartSize
		u.LeavePartsOnError = false
		u.MaxUploadParts = 10000
	})

	return &S3Struct{
		uploader: uploader,
	}, nil
}

var (
	S3  S3Interface
	err error
)

func init() {
	S3, err = NewS3Client()
	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[S3] Failed to initialize S3 client")
	}
}

func (s *S3Struct) Upload(file *multipart.FileHeader) (string, error) {
	// Validate input
	if file == nil {
		err := errors.New("file is nil")
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[S3][Upload] invalid input")
		return "", err
	}

	// Open file
	fileContent, err := file.Open()
	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[S3][Upload] failed to open file")
		return "", err
	}

	// Generate unique file name
	timeNow := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeNow, file.Filename)

	// Determine content type with fallback
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream" // Fallback for unknown file types
	}

	// Start Go routine to upload file
	go func() {
		defer func() {
			if cerr := fileContent.Close(); cerr != nil {
				log.Warn(log.LogInfo{
					"error": cerr.Error(),
				}, "[S3][Upload] failed to close file")
			}
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_, err := s.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
			Bucket:      aws.String(env.AppEnv.AWSS3BucketName),
			Key:         aws.String(fileName),
			Body:        fileContent,
			ACL:         aws.String("public-read"),
			ContentType: aws.String(contentType),
		})
		if err != nil {
			log.Error(log.LogInfo{
				"error":    err.Error(),
				"fileName": fileName,
			}, "[S3][Upload] failed to upload file")
		} else {
			log.Info(log.LogInfo{
				"fileName": fileName,
			}, "[S3][Upload] file uploaded successfully")
		}
	}()

	// Return the file name immediately
	return fmt.Sprintf("%s%s", env.AppEnv.AWSS3Path, fileName), nil
}
