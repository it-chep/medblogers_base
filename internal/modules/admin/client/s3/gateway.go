package s3

import (
	"context"
	"fmt"
	"io"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/freelancers/domain/doctor"
	"mime"
	"path/filepath"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/samber/lo"

	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . S3Client,S3PresignClient

type S3Client interface {
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

type S3PresignClient interface {
	PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

// Gateway клиент к S3
type Gateway struct {
	freelancersBucketName string // todo сделать общую структуру с баскетами
	doctorsBucketName     string
	blogsBucketName       string
	region                string
	client                S3Client
	presignClient         S3PresignClient
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
func NewGateway(freelancerBucketName, doctorsBucketName, blogsBucketName string, client S3Client, presignClient S3PresignClient) *Gateway {
	return &Gateway{
		client:                client,
		presignClient:         presignClient,
		freelancersBucketName: freelancerBucketName,
		doctorsBucketName:     doctorsBucketName,
		blogsBucketName:       blogsBucketName,
	}
}

// GeneratePresignedURL .
func (g *Gateway) GeneratePresignedURL(ctx context.Context, s3Key string) (string, error) {
	req, err := g.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(g.freelancersBucketName),
		Key:    aws.String(s3Key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Hour
	})
	if err != nil {
		return "", fmt.Errorf("ошибка при генерации presigned URL: %w", err)
	}

	return req.URL, nil
}

// PutBlogPhoto загружает фотографию в хранилище
func (g *Gateway) PutBlogPhoto(ctx context.Context, file io.Reader, filename string) (string, error) {
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	objectKey := fmt.Sprintf("images/%s", filename)

	// Загружаем файл в S3
	_, err := g.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(g.blogsBucketName),
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

// DelBlogPhoto удаляет фотографию из хранилища
func (g *Gateway) DelBlogPhoto(ctx context.Context, filename string) error {
	// Формируем полный ключ объекта
	objectKey := fmt.Sprintf("images/%s", filename)

	// Удаляем файл из S3
	_, err := g.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(g.blogsBucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func (g *Gateway) PutDoctorPhoto(ctx context.Context, file io.Reader, filename string) (string, error) {
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	objectKey := fmt.Sprintf("images/user_%s", filename)

	// Загружаем файл в S3
	_, err := g.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(g.doctorsBucketName),
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

func (g *Gateway) DelDoctorPhoto(ctx context.Context, filename string) error {
	// Формируем полный ключ объекта
	objectKey := fmt.Sprintf("images/user_%s", filename)

	// Удаляем файл из S3
	_, err := g.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(g.doctorsBucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// GetPhotoLink .
func (g *Gateway) GetPhotoLink(s3Key string) string {
	return fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", g.freelancersBucketName, s3Key)
}

// GetDoctorsPhotos .
func (g *Gateway) GetDoctorsPhotos(ctx context.Context) (map[doctor.S3Key]string, error) {
	filesMap := make(map[doctor.S3Key]string)
	paginator := s3.NewListObjectsV2Paginator(g.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(g.doctorsBucketName),
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

			filesMap[doctor.S3Key(key)] = fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", g.doctorsBucketName, key)
		}
	}

	return filesMap, nil
}
