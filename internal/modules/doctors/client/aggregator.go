package client

import (
	"medblogers_base/internal/modules/doctors/client/s3"
	"medblogers_base/internal/modules/doctors/client/salebot"
	"medblogers_base/internal/modules/doctors/client/subscribers"
)

type Aggregator struct {
	Subscribers *subscribers.Gateway
	S3          *s3.Gateway
	Salebot     *salebot.Gateway
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		Subscribers: subscribers.NewGateway(),
		S3:          s3.NewGateway(),
		Salebot:     salebot.NewGateway(),
	}
}
