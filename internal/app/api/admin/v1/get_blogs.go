package v1

import (
	"context"
	"github.com/google/uuid"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
	pkgctx "medblogers_base/internal/pkg/context"
)

func (i *Implementation) GetBlogs(ctx context.Context, req *desc.GetBlogsRequest) (resp *desc.GetBlogsResponse, _ error) {
	email := pkgctx.GetEmailFromContext(ctx)
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, email, "/api/v1/admin/blog", func(ctx context.Context) error {
		resp = &desc.GetBlogsResponse{}

		resp.BlogId = uuid.New().String()
		return nil
	})
}
