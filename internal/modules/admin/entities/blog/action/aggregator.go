package action

import (
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modu
	"medblogers_base/internal/modules/admin/entities/blog/action/create_draft_blog"
	"medblogers_base/internal/modules/admin/entities/blog/action/delete_blog_image"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blog_by_id"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blogs"
	"medblogers_base/internal/modules/admin/entities/blog/action/publish_blog"
	"medblogers_base/internal/modules/admin/entities/blog/action/save_blog_image"
	"medblogers_base/internal/modules/admin/entities/blog/action/unpublish_blog"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/postgres"
)

type BlogModuleAggregator struct {
	CreateDraftBlog *create_draft_blog.Action
	UpdateDraftBlog *update_draft_blog.Action

	SaveBlogImage   *save_blog_image.Action
	DeleteBlogImage *delete_blog_image.Action

	GetBlogs    *get_blogs.Action
	GetBlogByID *get_blog_by_id.Action

	PublishBlog   *publish_blog.Action
	UnPublishBlog *unpublish_blog.Action
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator) *BlogModuleAggregator {
	return &BlogModuleAggregator{
		CreateDraftBlog: create_draft_blog.New(pool),
		UpdateDraftBlog: update_draft_blog.New(pool),

		SaveBlogImage:   save_blog_image.New(pool, clients),
		DeleteBlogImage: delete_blog_image.New(pool, clients),

		GetBlogs:    get_blogs.New(pool),
		GetBlogByID: get_blog_by_id.New(pool),

		PublishBlog:   publish_blog.New(pool),
		UnPublishBlog: unpublish_blog.New(pool),
	}
}
