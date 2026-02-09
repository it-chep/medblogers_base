package dto

import (
	"github.com/samber/lo"
	"medblogers_base/internal/modules/admin/client/subscribers/indto"
)

type SubscribersData struct {
	Key       string
	SubsCount int64
}
type UpdateSubscribersRequest struct {
	Data []SubscribersData
}

func (r *UpdateSubscribersRequest) ToInDTO() []indto.UpdateSubscribersItem {
	return lo.Map(r.Data, func(item SubscribersData, _ int) indto.UpdateSubscribersItem {
		return indto.UpdateSubscribersItem{
			SubsCount: item.SubsCount,
			Key:       indto.NewSocialMedia(item.Key),
		}
	})
}
