package save_blog_image

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"medblogers_base/internal/modules/admin/action/blog/action/save_blog_image/dal"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/postgres"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type S3 interface {
	PutBlogPhoto(ctx context.Context, file io.Reader, filename string) (string, error)
}

type Action struct {
	s3Gateway S3
	dal       *dal.Repository
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator) *Action {
	return &Action{
		dal:       dal.NewRepository(pool),
		s3Gateway: clients.S3,
	}
}

func (a *Action) Do(ctx context.Context, blogID uuid.UUID, imageData []byte) (uuid.UUID, string, error) {
	imageID, _ := uuid.NewV7()

	reader := bytes.NewReader(imageData)
	contentType := http.DetectContentType(imageData)
	fileType := strings.Split(contentType, "/")[1]
	filename := fmt.Sprintf("%s.%s", imageID.String(), fileType)

	_, err := a.s3Gateway.PutBlogPhoto(ctx, reader, filename)
	if err != nil {
		return uuid.Nil, "", err
	}

	err = a.dal.SaveBlogPhoto(ctx, blogID, imageID, fileType)
	if err != nil {
		return uuid.Nil, "", err
	}

	// todo некрасиво поправить хардкод
	return imageID, fmt.Sprintf("https://storage.yandexcloud.net/medblogers-blogs/images/%s", filename), nil
}
