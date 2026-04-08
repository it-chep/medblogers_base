package blog

import (
	"fmt"

	"github.com/google/uuid"
)

// SetPrimaryPhotoURL сетим главную фотку
func (b *Blog) SetPrimaryPhotoURL(bucket string, photoID uuid.UUID, fileType string) {
	filename := fmt.Sprintf("%s.%s", photoID.String(), fileType)

	b.primaryPhotoURL = fmt.Sprintf("https://storage.yandexcloud.net/%s/images/%s", bucket, filename)
}

// SetViewsCount устанавливает количество просмотров статьи.
func (b *Blog) SetViewsCount(viewsCount int64) {
	b.viewsCount = viewsCount
}
