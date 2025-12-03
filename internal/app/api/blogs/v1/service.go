package v1

import (
	"medblogers_base/internal/modules/blogs"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
)

type Implementation struct {
	desc.UnimplementedBlogServiceServer

	blogs *blogs.Module
}

func NewService(module *blogs.Module) *Implementation {
	return &Implementation{
		blogs: module,
	}
}
