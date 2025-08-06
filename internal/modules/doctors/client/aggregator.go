package client

import (
	"fmt"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/doctors/client/s3"
	"medblogers_base/internal/modules/doctors/client/salebot"
	"medblogers_base/internal/modules/doctors/client/subscribers"
	pkgHttp "medblogers_base/internal/pkg/http"
)

type Aggregator struct {
	Subscribers *subscribers.Gateway
	S3          *s3.Gateway
	Salebot     *salebot.Gateway
}

func NewAggregator(httpConns map[string]pkgHttp.Executor, cfg config.AppConfig) *Aggregator {
	return &Aggregator{
		Subscribers: subscribers.NewGateway(
			fmt.Sprintf("%s:%s", cfg.GetSubscribersHost(), cfg.GetSubscribersPort()),
			httpConns[config.Subscribers],
		),
		S3: s3.NewGateway(cfg.GetUserPhotosBucket(), s3.NewS3Client(cfg.GetS3Config()), s3.NewPresignClient(cfg.GetS3Config())),
		Salebot: salebot.NewGateway(
			cfg.GetSalebotHost(),
			httpConns[config.Salebot],
		),
	}
}
