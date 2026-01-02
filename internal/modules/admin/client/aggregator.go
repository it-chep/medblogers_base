package client

import (
	"fmt"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/admin/client/s3"
	"medblogers_base/internal/modules/admin/client/subscribers"
	pkgHttp "medblogers_base/internal/pkg/http"
)

type Aggregator struct {
	S3          *s3.Gateway
	Subscribers *subscribers.Gateway
}

func NewAggregator(httpConns map[string]pkgHttp.Executor, cfg config.AppConfig) *Aggregator {
	return &Aggregator{
		S3: s3.NewGateway(
			cfg.GetFreelancersPhotosBucket(),
			cfg.GetUserPhotosBucket(),
			cfg.GetBlogsPhotosBucket(),
			s3.NewS3Client(cfg.GetS3Config()),
			s3.NewPresignClient(cfg.GetS3Config()),
		),
		Subscribers: subscribers.NewGateway(
			fmt.Sprintf("%s:%s", cfg.GetSubscribersHost(), cfg.GetSubscribersPort()),
			httpConns[config.Subscribers],
		),
	}
}
