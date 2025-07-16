package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"io"
	"medblogers_base/internal/config"
	"mime"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Gateway клиент к S3
type Gateway struct {
	bucketName    string
	client        *s3.Client
	presignClient *s3.PresignClient
}

// NewGateway .
func NewGateway(bucketName string, cfg config.S3Config) *Gateway {
	s3Cfg, err := s3config.LoadDefaultConfig(
		context.Background(),
		s3config.WithRegion(cfg.Region),
		s3config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKey, cfg.SecretKey, "",
		)),
		s3config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           os.Getenv(cfg.Endpoint),
					SigningRegion: os.Getenv(cfg.Region),
				}, nil
			}),
		),
	)
	if err != nil {
		return nil
	}

	client := s3.NewFromConfig(s3Cfg)
	presignClient := s3.NewPresignClient(client)

	return &Gateway{
		client:        client,
		presignClient: presignClient,
		bucketName:    bucketName,
	}
}

// GetUserPhotos получение фотографий врачей
func (g *Gateway) GetUserPhotos(ctx context.Context) error {
	return nil
}

// GeneratePresignedURL .
func (g *Gateway) GeneratePresignedURL(ctx context.Context, s3Key string) (string, error) {
	req, err := g.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(g.bucketName),
		Key:    aws.String(s3Key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Hour
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return req.URL, nil
}

// PutObject загружает фотографию врача в хранилище
func (g *Gateway) PutObject(ctx context.Context, file io.Reader, filename string) (string, error) {
	// Генерируем уникальный ключ для файла
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	// todo
	//objectKey := fmt.Sprintf("doctors/%s%s", uuid.New().String(), ext)

	// Загружаем файл в S3
	_, err := g.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(g.bucketName),
		Key:         aws.String(""), // todo сделать
		Body:        file,
		ContentType: aws.String(contentType),
		Metadata: map[string]string{
			"uploaded_at": time.Now().Format(time.RFC3339),
			"origin_name": filename,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return "objectKey", nil
}
