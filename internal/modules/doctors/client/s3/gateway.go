package s3

import (
	"context"
)

// https://yandex.cloud/ru/docs/storage/tools/aws-sdk-go

type Gateway struct {
}

func NewGateway() *Gateway {
	return &Gateway{}
}

func (g *Gateway) GetUserPhotos(ctx context.Context) error {
	return nil
}

func (g *Gateway) GeneratePresignedURL(ctx context.Context) error {
	return nil
}
