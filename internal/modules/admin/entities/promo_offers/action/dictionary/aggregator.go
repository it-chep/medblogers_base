package dictionary

import (
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/create_content_format"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/create_cooperation_type"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/create_topic"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/get_content_formats"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/get_cooperation_types"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/get_social_networks"
	"medblogers_base/internal/modules/admin/entities/promo_offers/action/dictionary/get_topics"
	"medblogers_base/internal/pkg/postgres"
)

type PromoOfferDictionaryAggregator struct {
	CreateTopic           *create_topic.Action
	CreateCooperationType *create_cooperation_type.Action
	CreateContentFormat   *create_content_format.Action
	GetTopics             *get_topics.Action
	GetCooperationTypes   *get_cooperation_types.Action
	GetContentFormats     *get_content_formats.Action
	GetSocialNetworks     *get_social_networks.Action
}

func New(pool postgres.PoolWrapper) *PromoOfferDictionaryAggregator {
	return &PromoOfferDictionaryAggregator{
		CreateTopic:           create_topic.New(pool),
		CreateCooperationType: create_cooperation_type.New(pool),
		CreateContentFormat:   create_content_format.New(pool),
		GetTopics:             get_topics.New(pool),
		GetCooperationTypes:   get_cooperation_types.New(pool),
		GetContentFormats:     get_content_formats.New(pool),
		GetSocialNetworks:     get_social_networks.New(pool),
	}
}
