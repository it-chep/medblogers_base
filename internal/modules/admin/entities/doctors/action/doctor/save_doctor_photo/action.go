package save_doctor_photo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/save_doctor_photo/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/save_doctor_photo/service/image"
	commondal "medblogers_base/internal/modules/admin/entities/doctors/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"
	"net/http"
	"strings"
)

type ActionDal interface {
	SaveDoctorImage(ctx context.Context, doctorID int64, image string) error
}

type CommonDal interface {
	GetDoctorByID(ctx context.Context, doctorID int64) (*doctor.Doctor, error)
}

// Action активация доктора
type Action struct {
	commonDal CommonDal
	actionDal ActionDal
	image     *image.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: commondal.NewRepository(pool),
		actionDal: dal.NewRepository(pool),
		image:     image.New(clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64, imageData []byte) (string, error) {
	doc, err := a.commonDal.GetDoctorByID(ctx, doctorID)
	if err != nil {
		return "", err
	}

	filename, err := a.generateNewImageName(doc, imageData)
	if err != nil {
		return "", err
	}
	newImageURL, err := a.image.SetImage(ctx, bytes.NewReader(imageData), filename)
	if err != nil {
		return "", err
	}

	err = a.actionDal.SaveDoctorImage(ctx, doctorID, newImageURL)
	if err != nil {
		return "", err
	}

	// Если раньше была фотография, то мы должны ее удалить
	if len(doc.GetS3Key().String()) != 0 {
		err = a.image.DeleteImage(ctx, doc.GetS3Key().String())
		if err != nil {
			return "", err
		}
	}

	return a.image.GetImageURL(newImageURL), nil
}

func (a *Action) generateNewImageName(doc *doctor.Doctor, imageData []byte) (string, error) {
	randUUID, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(imageData)
	fileType := strings.Split(contentType, "/")[1]
	filename := fmt.Sprintf("%s_%s.%s", doc.GetSlug(), randUUID, fileType)

	return filename, nil
}
