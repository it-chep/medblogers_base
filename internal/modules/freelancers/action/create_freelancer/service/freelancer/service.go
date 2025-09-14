package freelancer

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/slug"
	"strings"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	CreateFreelancer(ctx context.Context, createDTO dto.CreateRequest) (int64, error)
	CreateAdditionalCities(ctx context.Context, medblogersID int64, citiesIDs []int64) error
	CreateAdditionalSpecialities(ctx context.Context, medblogersID int64, specialitiesIDs []int64) error
	CreateSocialNetworks(ctx context.Context, freelancerID int64, networkIDs []int64) error
	CreatePriceList(ctx context.Context, freelancerID int64, priceList dto.PriceList) error
}

// todo сделать номрально
var priceCategoriesMap = map[int64]int64{
	1: 1000,
	2: 20000,
	3: 50000,
	4: 100_000,
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateOrUpdate(ctx context.Context, createDTO dto.CreateRequest) (dto.CreateRequest, error) {
	logger.Message(ctx, "[Create] Создание слага и имени")

	createDTO.Name = s.createName(createDTO.LastName, createDTO.FirstName, createDTO.MiddleName)
	createDTO.Slug = slug.New(createDTO.Name)
	createDTO.PriceCategory = s.definePriceCategory(createDTO.PriceList)

	logger.Message(ctx, "[Create] Сохранение фрилансера в базе")

	medblogersID, err := s.storage.CreateFreelancer(ctx, createDTO)
	if err != nil {
		return dto.CreateRequest{}, err
	}

	createDTO.ID = medblogersID
	createDTO.AdditionalCities = append(createDTO.AdditionalCities, createDTO.MainCityID)
	createDTO.AdditionalSpecialties = append(createDTO.AdditionalSpecialties, createDTO.MainSpecialityID)

	// todo транзакция
	g := async.NewGroup()

	logger.Message(ctx, "[Create] Сохранение дополнительных параметров")
	g.Go(func() {
		err = s.storage.CreateAdditionalCities(ctx, medblogersID, createDTO.AdditionalCities)
		if err != nil {
			logger.Error(ctx, "[Create] Ошибка при сохранении доп городов ", err)
		}
	})
	g.Go(func() {
		err = s.storage.CreateAdditionalSpecialities(ctx, medblogersID, createDTO.AdditionalSpecialties)
		if err != nil {
			logger.Error(ctx, "[Create] Ошибка при сохранении доп специальностей ", err)
		}
	})
	g.Go(func() {
		err = s.storage.CreateSocialNetworks(ctx, medblogersID, createDTO.SocialNetworks)
		if err != nil {
			logger.Error(ctx, "[Create] Ошибка при сохранении соц сетей ", err)
		}
	})
	g.Go(func() {
		err = s.storage.CreatePriceList(ctx, medblogersID, createDTO.PriceList)
		if err != nil {
			logger.Error(ctx, "[Create] Ошибка при сохранении прайс листа", err)
		}
	})

	g.Wait()

	return createDTO, nil
}

func (s *Service) createName(lastName, firstName, middleName string) string {
	parts := []string{
		strings.TrimSpace(lastName),
		strings.TrimSpace(firstName),
		strings.TrimSpace(middleName),
	}

	// Удаляем пустые части
	var nonEmptyParts []string
	for _, part := range parts {
		if part != "" {
			nonEmptyParts = append(nonEmptyParts, part)
		}
	}

	return strings.Join(nonEmptyParts, " ")
}

func (s *Service) definePriceCategory(priceList dto.PriceList) int64 {
	if len(priceList) == 0 {
		return 1
	}

	// Считаем общую сумму цен
	var totalSum int64 = 0
	for _, item := range priceList {
		totalSum += item.Price
	}

	// Вычисляем среднюю цену
	averagePrice := totalSum / int64(len(priceList))

	// Определяем категорию на основе средней цены
	var category int64 = 1

	// Проходим по категориям от самой низкой к самой высокой
	for cat, maxPrice := range priceCategoriesMap {
		if averagePrice <= maxPrice {
			category = cat
			break
		}
		// Если цена превышает все категории, останется последняя проверенная
	}

	return category
}
