package doctor_detail

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/dal"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/service/additional_items"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/service/doctor"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/service/enricher"
	"medblogers_base/internal/modules/doctors/client"
)

// Action получение детальной информации о докторе
type Action struct {
	doctorService          *doctor.Service
	additionalItemsService *additional_items.Service
	subscribersEnricher    *enricher.SubscribersEnricher
	imagesEnricher         *enricher.ImageEnricher
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	repo := dal.NewRepository(pool)
	return &Action{
		doctorService:          doctor.New(doctor_dal.NewRepository(pool)),
		additionalItemsService: additional_items.New(repo),
		subscribersEnricher:    enricher.NewSubscribersEnricher(clients.Subscribers),
		imagesEnricher:         enricher.NewImageEnricher(clients.S3),
	}
}

func (a Action) Do(ctx context.Context, slug string) (*dto.DoctorDTO, error) {
	logger.Message(ctx, fmt.Sprintf("[DoctorDetail] Получение данных о докторе %s", slug))
	doc, err := a.doctorService.GetDoctorInfo(ctx, slug)
	if err != nil {
		logger.Error(ctx, "Ошибка при получении доктора", err)
		return &dto.DoctorDTO{}, err // 404 not found
	}

	docDTO := dto.New(doc)
	g := async.NewGroup()

	// получаем подписчиков
	g.Go(func() {
		docDTO, err = a.subscribersEnricher.Enrich(ctx, doc.GetID(), docDTO)
		if err != nil {
			logger.Error(ctx, "Ошибка при обогащении доктора подписчиками", err)
		}
	})

	// получаем доп города
	g.Go(func() {
		mainCity, cities, err := a.additionalItemsService.GetAdditionalCities(ctx, doc.GetID(), doc.GetMainCityID())
		if err != nil {
			logger.Error(ctx, "Ошибка при получении дополнительных городов", err)
		}
		// при ошибке будет пустой список
		docDTO.Cities = cities
		docDTO.MainCity = mainCity
	})

	// получаем доп специальности
	g.Go(func() {
		mainSpeciality, specialities, err := a.additionalItemsService.GetAdditionalSpecialities(ctx, doc.GetID(), doc.GetMainSpecialityID())
		if err != nil {
			logger.Error(ctx, "Ошибка при получении дополнительных специальностей", err)
		}
		// при ошибке будет пустой список
		docDTO.Specialities = specialities
		docDTO.MainSpeciality = mainSpeciality
	})

	// обогащение фоткой с S3
	g.Go(func() {
		image, err := a.imagesEnricher.Enrich(ctx, doc.GetS3Key().String())
		if err != nil {
			image = "https://storage.yandexcloud.net/medblogers-photos/images/zag.png"
			logger.Error(ctx, "Ошибка при получении дополнительных специальностей", err)
		}
		docDTO.Image = image
	})

	g.Wait()

	return docDTO, nil
}
