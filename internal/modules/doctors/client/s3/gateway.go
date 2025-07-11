package s3

import (
	"context"
)

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
