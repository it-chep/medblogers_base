package freelancer_detail

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/freelancers/action/freelancer_detail/dal"
	"medblogers_base/internal/modules/freelancers/action/freelancer_detail/dto"
	"medblogers_base/internal/modules/freelancers/action/freelancer_detail/service/additional_items"
	"medblogers_base/internal/modules/freelancers/action/freelancer_detail/service/enricher"
	"medblogers_base/internal/modules/freelancers/action/freelancer_detail/service/freelancer"
	"medblogers_base/internal/modules/freelancers/client"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение настроек главной страницы
type Action struct {
	freelancer        *freelancer.Service
	imageEnricher     *enricher.ImageEnricher
	additionalService *additional_items.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	repository := dal.NewRepository(pool)
	return &Action{
		freelancer:        freelancer.New(repository),
		imageEnricher:     enricher.NewImageEnricher(clients.S3),
		additionalService: additional_items.New(repository),
	}
}

func (a *Action) Do(ctx context.Context, slug string) (*dto.FreelancerDTO, error) {
	logger.Message(ctx, fmt.Sprintf("[FreelancerDetail] Получение данных о фрилансере %s", slug))
	frlncr, err := a.freelancer.GetFreelancerInfo(ctx, slug)
	if err != nil {
		logger.Error(ctx, "Ошибка при получении фрилансера", err)
		return &dto.FreelancerDTO{}, err // 404 not found
	}

	frlncrDTO := dto.New(frlncr)
	g := async.NewGroup()

	// получаем доп города
	g.Go(func() {
		mainCity, cities, gErr := a.additionalService.GetAdditionalCities(ctx, frlncr.GetID(), frlncr.GetMainCityID())
		if gErr != nil {
			logger.Error(ctx, "Ошибка при получении дополнительных городов", gErr)
		}
		// при ошибке будет пустой список
		frlncrDTO.Cities = cities
		frlncrDTO.MainCity = mainCity
	})

	// получаем доп специальности
	g.Go(func() {
		mainSpeciality, specialities, gErr := a.additionalService.GetAdditionalSpecialities(ctx, frlncr.GetID(), frlncr.GetMainSpecialityID())
		if gErr != nil {
			logger.Error(ctx, "Ошибка при получении дополнительных специальностей", gErr)
		}
		// при ошибке будет пустой список
		frlncrDTO.Specialities = specialities
		frlncrDTO.MainSpeciality = mainSpeciality
	})

	// прайс лист
	g.Go(func() {
		priceList, gErr := a.additionalService.GetPriceList(ctx, frlncr.GetID())
		if gErr != nil {
			logger.Error(ctx, "Ошибка при получении прайс листа", gErr)
		}
		// при ошибке будет пустой список
		frlncrDTO.PriceList = priceList
	})

	// соц сети
	g.Go(func() {
		socialNetworks, gErr := a.additionalService.GetSocialNetworks(ctx, frlncr.GetID())
		if gErr != nil {
			logger.Error(ctx, "Ошибка при получении соц сетей", gErr)
		}
		// при ошибке будет пустой список
		frlncrDTO.SocialNetworks = socialNetworks
	})

	// обогащение фоткой с S3
	g.Go(func() {
		image, gErr := a.imageEnricher.Enrich(ctx, frlncr.GetS3Image())
		if gErr != nil {
			image = "https://storage.yandexcloud.net/medblogers-photos/images/zag.png"
			logger.Error(ctx, "Ошибка при получении дополнительных специальностей", gErr)
		}
		frlncrDTO.Image = image
	})

	g.Wait()

	return frlncrDTO, nil
}
