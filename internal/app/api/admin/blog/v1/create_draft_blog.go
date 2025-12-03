package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
)

func (i *Implementation) CreateDraftBlog(ctx context.Context, req *desc.CreateDraftBlogRequest) (resp *desc.CreateDraftBlogResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/create", func(ctx context.Context) error {
		resp = &desc.CreateDraftBlogResponse{}

		blogID, err := i.admin.Actions.BlogModule.CreateDraftBlog.Do(ctx, req.GetTitle())
		if err != nil {
			return err
		}

		resp.BlogId = blogID.String()

		return nil
	})
}
