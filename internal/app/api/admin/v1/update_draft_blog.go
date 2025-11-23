package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
	pkgctx "medblogers_base/internal/pkg/context"
)

func (i *Implementation) UpdateDraftBlog(ctx context.Context, req *desc.UpdateDraftBlogRequest) (resp *desc.UpdateDraftBlogResponse, _ error) {
	email := pkgctx.GetEmailFromContext(ctx)
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, email, "/api/v1/auth/blog/{id}/update", func(ctx context.Context) error {
		resp = &desc.UpdateDraftBlogResponse{}

		return nil
	})
}
