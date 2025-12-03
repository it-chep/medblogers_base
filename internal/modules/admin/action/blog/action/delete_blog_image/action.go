package delete_blog_image

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/admin/action/blog/action/delete_blog_image/dal"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

type S3 interface {
	DelBlogPhoto(ctx context.Context, filename string) error
}

type Action struct {
	dal       *dal.Repository
	s3Gateway S3
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator) *Action {
	return &Action{
		dal:       dal.NewRepository(pool),
		s3Gateway: clients.S3,
	}
}

// Do .
func (a *Action) Do(ctx context.Context, blogID, imageID uuid.UUID) error {
	imageDTO, err := a.dal.GetBlogImageByID(ctx, imageID)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.%s", imageID.String(), imageDTO.FileType)
	err = a.s3Gateway.DelBlogPhoto(ctx, fileName)
	if err != nil {
		return err
	}

	err = a.dal.DeleteBlogImage(ctx, blogID, imageID)
	if err != nil {
		return err
	}

	return nil
}
