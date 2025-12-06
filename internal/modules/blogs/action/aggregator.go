package action

import (
	"medblogers_base/internal/modules/blogs/action/get_blog_detail"
	"medblogers_base/internal/modules/blogs/action/get_blogs"
	"medblogers_base/internal/modules/blogs/action/get_top_blogs"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	GetBlogDetail *get_blog_detail.Action
	GetBlogs      *get_blogs.Action
	GetTopBlogs   *get_top_blogs.Action
}

// NewAggregator конструктор
func NewAggregator(pool postgres.PoolWrapper) *Aggregator {
	return &Aggregator{
		GetBlogs:      get_blogs.New(pool),
		GetTopBlogs:   get_top_blogs.New(pool),
		GetBlogDetail: get_blog_detail.New(pool),
	}
}
