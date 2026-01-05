package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) ChangeRecommendationSort(ctx context.Context, req *desc.ChangeRecommendationSortRequest) (resp *desc.ChangeRecommendationSortResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/change_recommendation_sort/{id}", func(ctx context.Context) error {
		return nil
	})
}
