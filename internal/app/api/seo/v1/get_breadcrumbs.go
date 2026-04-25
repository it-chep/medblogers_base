package v1

import (
	"context"
	seoDTO "medblogers_base/internal/modules/seo/action/get_breadcrumbs/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"

	"github.com/samber/lo"
)

// GetBreadcrumbs ...
func (i *Implementation) GetBreadcrumbs(ctx context.Context, req *desc.GetBreadcrumbsRequest) (resp *desc.GetBreadcrumbsResponse, err error) {
	breadcrumbs, err := i.seo.Actions.GetBreadcrumbs.Do(ctx, req.GetPath())
	if err != nil {
		return nil, err
	}

	return &desc.GetBreadcrumbsResponse{
		Breadcrumbs: lo.Map(breadcrumbs, func(item seoDTO.Breadcrumb, _ int) *desc.GetBreadcrumbsResponse_Breadcrumb {
			return &desc.GetBreadcrumbsResponse_Breadcrumb{
				Name: item.Name,
				Path: item.Path,
			}
		}),
	}, nil
}
