package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"

	"github.com/samber/lo"
)

func (i *Implementation) GetSitemapInfo(ctx context.Context, req *desc.GetSitemapInfoRequest) (resp *desc.GetSitemapInfoResponse, err error) {
	urls, err := i.seo.Actions.GetSitemapInfo.Do(ctx)
	if err != nil {
		return nil, err
	}

	resp = &desc.GetSitemapInfoResponse{
		Items: lo.Map(urls, func(url string, index int) *desc.GetSitemapInfoResponse_SitemapItem {
			return &desc.GetSitemapInfoResponse_SitemapItem{
				Url: url,
			}
		}),
	}

	return resp, nil
}
