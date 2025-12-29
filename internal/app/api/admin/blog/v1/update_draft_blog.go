package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/blog/action/update_draft_blog/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"

	"github.com/google/uuid"
)

func (i *Implementation) UpdateDraftBlog(ctx context.Context, req *desc.UpdateDraftBlogRequest) (resp *desc.UpdateDraftBlogResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/update", func(ctx context.Context) error {
		resp = &desc.UpdateDraftBlogResponse{}

		updateRequest := dto.Request{
			Name:               req.GetTitle(),
			Slug:               req.GetSlug(),
			Body:               req.GetBody(),
			IsActive:           req.GetIsActive(),
			PreviewText:        req.GetPreviewText(),
			SocietyPreviewText: req.GetSocietyPreview(),
			AdditionalSEOText:  req.GetAdditionalSeoText(),
			OrderingNumber:     req.GetOrderingNumber(),
		}

		err := i.admin.Actions.BlogModule.UpdateDraftBlog.Do(ctx, uuid.MustParse(req.GetBlogId()), updateRequest)
		if err != nil {
			return err
		}

		return nil
	})
}
