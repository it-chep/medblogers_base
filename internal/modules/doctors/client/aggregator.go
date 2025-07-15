package client

import (
	"fmt"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/doctors/client/s3"
	"medblogers_base/internal/modules/doctors/client/salebot"
	"medblogers_base/internal/modules/doctors/client/subscribers"
	"net/http"
	"time"
)

type Aggregator struct {
	Subscribers *subscribers.Gateway
	S3          *s3.Gateway
	Salebot     *salebot.Gateway
}

func NewAggregator(config *config.Config) *Aggregator {
	return &Aggregator{
		Subscribers: subscribers.NewGateway(
			fmt.Sprintf("%s:%s", config.SubscribersClient.Host, config.SubscribersClient.Port),
			&http.Client{
				Timeout: 3 * time.Second,
			}),
		S3: s3.NewGateway(),
		Salebot: salebot.NewGateway(&http.Client{
			Timeout: 3 * time.Second,
		}),
	}
}
