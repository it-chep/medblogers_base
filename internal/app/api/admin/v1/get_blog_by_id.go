package v1

import (
	"context"
	"github.com/google/uuid"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
)

func (i *Implementation) GetBlogByID(ctx context.Context, req *desc.GetBlogByIDRequest) (resp *desc.GetBlogByIDResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}", func(ctx context.Context) error {
		resp = &desc.GetBlogByIDResponse{}

		_, err := i.admin.Actions.BlogModule.GetBlogByID.Do(ctx, uuid.MustParse(req.GetBlogId()))
		if err != nil {
			return err
		}

		return nil
	})
}
