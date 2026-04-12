package dto

import "medblogers_base/internal/modules/blogs/domain/blog"

type Response struct {
	Blogs blog.Blogs
}
