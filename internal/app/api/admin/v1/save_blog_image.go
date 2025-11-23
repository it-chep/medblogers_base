package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
	pkgctx "medblogers_base/internal/pkg/context"
)

func (i *Implementation) SaveBlogImage(ctx context.Context, req *desc.SaveBlogImageRequest) (*desc.SaveBlogImageResponse, error) {
	email := pkgctx.GetEmailFromContext(ctx)
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return nil, executor(ctx, email, "/api/v1/auth/blog/{id}/save_image", func(ctx context.Context) error {
		return nil
	})
}
