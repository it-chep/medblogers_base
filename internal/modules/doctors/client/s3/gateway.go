package s3

import (
	"context"
	"fmt"
	"io"
	"medblogers_base/internal/config"
	"medblogers_base/internal/pkg/logger"
	"mime"
	"path/filepath"
	"strings"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/samber/lo"

	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . S3Client,S3PresignClient

// S3Client .
type S3Client interface {
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// S3PresignClient .
type S3PresignClient interface {
	PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

// Gateway клиент к S3
type Gateway struct {
	bucketName    string
	client        S3Client
	presignClient S3PresignClient
}

func NewS3Client(cfg config.S3Config) S3Client {
	s3Cfg, err := s3config.LoadDefaultConfig(context.Background(),
		s3config.WithRegion(cfg.Region),
		s3config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKey, cfg.SecretKey, "",
		)),
		s3config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           cfg.Endpoint,
					SigningRegion: cfg.Region,
				}, nil
			})),
	)
	if err != nil {
		return nil
	}

	return s3.NewFromConfig(s3Cfg)
}

func NewPresignClient(cfg config.S3Config) S3PresignClient {
	s3Cfg, err := s3config.LoadDefaultConfig(context.Background(),
		s3config.WithRegion(cfg.Region),
		s3config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKey, cfg.SecretKey, "",
		)),
		s3config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           cfg.Endpoint,
					SigningRegion: cfg.Region,
				}, nil
			})),
	)
	if err != nil {
		return nil
	}

	client := s3.NewFromConfig(s3Cfg)
	return s3.NewPresignClient(client)
}

// NewGateway .
func NewGateway(bucketName string, client S3Client, presignClient S3PresignClient) *Gateway {
	return &Gateway{
		client:        client,
		presignClient: presignClient,
		bucketName:    bucketName,
	}
}

// todo закрыть это место кешом, чтобы не было постоянных проходок

// GetUserPhotos получение фотографий врачей из Yandex Object Storage
func (g *Gateway) GetUserPhotos(ctx context.Context) (map[string]string, error) {
	logger.Message(ctx, "[S3] Получение фотографий пользователей из Yandex Storage")

	filesMap := make(map[string]string)
	paginator := s3.NewListObjectsV2Paginator(g.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(g.bucketName),
		Prefix: aws.String("images/user_"),
	})

	// Обрабатываем все страницы результатов
	for paginator.HasMorePages() {
		resp, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list objects: %w", err)
		}

		// Обрабатываем объекты на текущей странице
		for _, object := range resp.Contents {
			key := aws.ToString(object.Key)
			if key == "" {
				continue
			}

			parts := strings.Split(key, "_")
			if len(parts) < 2 {
				continue
			}
			slug := parts[1]

			filesMap[slug] = fmt.Sprintf("https://storage.yandexcloud.net/%s/%s",
				g.bucketName, key)
		}
	}

	return filesMap, nil
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
		return "", fmt.Errorf("ошибка при генерации presigned URL: %w", err)
	}

	return req.URL, nil
}

// GetPhotoLink .
func (g *Gateway) GetPhotoLink(s3Key string) string {
	return fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", g.bucketName, s3Key)
}

// PutObject загружает фотографию врача в хранилище
func (g *Gateway) PutObject(ctx context.Context, file io.Reader, filename, slug string) (string, error) {
	// Генерируем уникальный ключ для файла
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	objectKey := fmt.Sprintf("images/user_%s_%s", slug, filename)

	// Загружаем файл в S3
	_, err := g.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(g.bucketName),
		Key:         lo.ToPtr(objectKey),
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

	return objectKey, nil
}
