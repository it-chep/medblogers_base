package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
	pkgctx "medblogers_base/internal/pkg/context"
)

func (i *Implementation) GetBlogByID(ctx context.Context, req *desc.GetBlogByIDRequest) (resp *desc.GetBlogByIDResponse, _ error) {
	email := pkgctx.GetEmailFromContext(ctx)
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, email, "/api/v1/admin/blog/{id}", func(ctx context.Context) error {
		resp = &desc.GetBlogByIDResponse{}

		return nil
	})
}
