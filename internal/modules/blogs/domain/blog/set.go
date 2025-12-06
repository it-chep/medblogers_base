package blog

import (
	"fmt"

	"github.com/google/uuid"
)

// SetPrimaryPhotoURL сетим главную фотку
func (b *Blog) SetPrimaryPhotoURL(photoID uuid.UUID, fileType string) {
	filename := fmt.Sprintf("%s.%s", photoID.String(), fileType)

	b.primaryPhotoURL = fmt.Sprintf("https://storage.yandexcloud.net/medblogers-blogs/images/%s", filename)
}
