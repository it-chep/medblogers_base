package dto

import (
	"medblogers_base/internal/modules/blogs/domain/blog"
	"medblogers_base/internal/modules/blogs/domain/category"
)

type Blog struct {
	*blog.Blog
	Categories category.Categories
}

type Response struct {
	Blogs []Blog
}
