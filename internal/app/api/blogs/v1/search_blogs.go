package v1

import (
	"context"
	"strconv"
	"strings"

	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) SearchBlogs(ctx context.Context, req *desc.SearchBlogsRequest) (*desc.SearchBlogsResponse, error) {
	query := strings.TrimSpace(req.GetQuery())
	if query == "" {
		return nil, status.Error(codes.InvalidArgument, "query is required")
	}

	resp, err := i.blogs.Actions.SearchBlogs.Do(ctx, query)
	if err != nil {
		return nil, err
	}

	searchItems := make([]*desc.SearchBlogsResponse_SearchItem, 0, len(resp.Blogs))
	for _, item := range resp.Blogs {
		searchItems = append(searchItems, &desc.SearchBlogsResponse_SearchItem{
			Title:       item.GetTitle(),
			Slug:        item.GetSlug(),
			PreviewText: item.GetPreviewText(),
			PhotoLink:   item.GetPrimaryPhotoURL(),
			ViewsCount:  strconv.FormatInt(item.GetViewsCount(), 10),
		})
	}

	return &desc.SearchBlogsResponse{
		Blogs: searchItems,
	}, nil
}
