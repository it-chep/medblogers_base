package client

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/freelancers/client/s3"
)

type Aggregator struct {
	S3 *s3.Gateway
}

func NewAggregator(cfg config.AppConfig) *Aggregator {
	return &Aggregator{
		S3: s3.NewGateway(
			cfg.GetUserPhotosBucket(),
			s3.NewS3Client(cfg.GetS3Config()),
			s3.NewPresignClient(cfg.GetS3Config()),
		),
	}
}
