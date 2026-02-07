package save_freelancer_photo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/save_freelancer_photo/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/save_freelancer_photo/service/image"
	commondal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/postgres"
	"net/http"
	"strings"
)

type ActionDal interface {
	SaveFreelancerImage(ctx context.Context, freelancerID int64, image string) error
}

type CommonDal interface {
	GetFreelancerByID(ctx context.Context, freelancerID int64) (*freelancer.Freelancer, error)
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

func (a *Action) Do(ctx context.Context, freelancerID int64, imageData []byte) (string, error) {
	frncer, err := a.commonDal.GetFreelancerByID(ctx, freelancerID)
	if err != nil {
		return "", err
	}

	filename, err := a.generateNewImageName(frncer, imageData)
	if err != nil {
		return "", err
	}
	newImageURL, err := a.image.SetImage(ctx, bytes.NewReader(imageData), filename)
	if err != nil {
		return "", err
	}

	err = a.actionDal.SaveFreelancerImage(ctx, freelancerID, newImageURL)
	if err != nil {
		return "", err
	}

	if len(frncer.GetS3Image()) != 0 {
		err = a.image.DeleteImage(ctx, frncer.GetS3Image())
		if err != nil {
			return "", err
		}
	}

	return a.image.GetImageURL(newImageURL), nil
}

func (a *Action) generateNewImageName(frncer *freelancer.Freelancer, imageData []byte) (string, error) {
	randUUID, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(imageData)
	fileType := strings.Split(contentType, "/")[1]
	filename := fmt.Sprintf("%s_%s.%s", frncer.GetSlug(), randUUID, fileType)

	return filename, nil
}
