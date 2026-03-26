package s3

import (
	"context"
	"fmt"
	"medblogers_base/internal/config"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Gateway struct {
	bucketName string
	client     S3Client
}

type S3Client interface {
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
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

func NewGateway(bucketName string, client S3Client) *Gateway {
	return &Gateway{
		bucketName: bucketName,
		client:     client,
	}
}

func (g *Gateway) GetBrandPhotoLink(s3Key string) string {
	if s3Key == "" {
		return ""
	}

	if strings.HasPrefix(s3Key, "http://") || strings.HasPrefix(s3Key, "https://") {
		return s3Key
	}

	return fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", g.bucketName, s3Key)
}

func (g *Gateway) GetBrandsPhotos(ctx context.Context) (map[string]string, error) {
	result := make(map[string]string)
	paginator := s3.NewListObjectsV2Paginator(g.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(g.bucketName),
		Prefix: aws.String("brands_photos/"),
	})

	for paginator.HasMorePages() {
		resp, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list brand photos: %w", err)
		}

		for _, object := range resp.Contents {
			key := aws.ToString(object.Key)

			if _, ok := result[key]; ok {
				continue
			}

			result[key] = g.GetBrandPhotoLink(key)
		}
	}

	return result, nil
}
