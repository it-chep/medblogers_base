package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"time"
)

func (i *Implementation) GetBlogDetail(ctx context.Context, req *desc.GetBlogDetailRequest) (*desc.GetBlogDetailResponse, error) {
	blog, err := i.blogs.Actions.GetBlogDetail.Do(ctx, req.GetBlogSlug())
	if err != nil {
		return nil, err
	}

	return &desc.GetBlogDetailResponse{
		Title:             blog.GetTitle(),
		Slug:              blog.GetSlug(),
		Body:              blog.GetBody(),
		PreviewText:       blog.GetPreviewText(),
		SocietyPreview:    blog.GetSocietyPreviewText(),
		AdditionalSeoText: blog.GetAdditionalSEOText(),
		CreatedAt:         blog.GetCreatedAt().Format(time.RFC3339),
		PhotoLink:         blog.GetPrimaryPhotoURL(),
	}, nil
}
