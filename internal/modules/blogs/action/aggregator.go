package action

import (
	"medblogers_base/internal/modules/blogs/action/filter_blogs"
	"medblogers_base/internal/modules/blogs/action/get_blog_detail"
	"medblogers_base/internal/modules/blogs/action/get_blogs"
	"medblogers_base/internal/modules/blogs/action/get_blogs_categories"
	"medblogers_base/internal/modules/blogs/action/get_doctor_blogs"
	"medblogers_base/internal/modules/blogs/action/get_top_blogs"
	"medblogers_base/internal/modules/blogs/client"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	GetBlogDetail      *get_blog_detail.Action
	GetBlogs           *get_blogs.Action
	GetTopBlogs        *get_top_blogs.Action
	GetDoctorBlogs     *get_doctor_blogs.Action
	GetBlogsCategories *get_blogs_categories.Action
	FilterBlogs        *filter_blogs.Action
}

// NewAggregator конструктор
func NewAggregator(pool postgres.PoolWrapper, clients *client.Aggregator) *Aggregator {
	return &Aggregator{
		GetBlogs:           get_blogs.New(pool),
		GetTopBlogs:        get_top_blogs.New(pool),
		GetBlogDetail:      get_blog_detail.New(pool, clients),
		GetDoctorBlogs:     get_doctor_blogs.New(pool),
		GetBlogsCategories: get_blogs_categories.New(pool),
		FilterBlogs:        filter_blogs.New(pool),
	}
}
