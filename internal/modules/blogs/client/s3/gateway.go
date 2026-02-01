package s3

import (
	"context"
	"fmt"
	"medblogers_base/internal/config"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . S3PresignClient

// S3PresignClient .
type S3PresignClient interface {
	PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

// Gateway клиент к S3
type Gateway struct {
	doctorBucket  string
	presignClient S3PresignClient
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
func NewGateway(doctorBucket string, presignClient S3PresignClient) *Gateway {
	return &Gateway{
		presignClient: presignClient,
		doctorBucket:  doctorBucket,
	}
}

// GetPhotoLink .
func (g *Gateway) GetPhotoLink(s3Key string) string {
	return fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", g.doctorBucket, s3Key)
}
