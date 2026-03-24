package filter_settings

import (
	"context"

	"github.com/samber/lo"

	actionDal "medblogers_base/internal/modules/promo_offers/action/filter_settings/dal"
	"medblogers_base/internal/modules/promo_offers/action/filter_settings/dto"
	commonDal "medblogers_base/internal/modules/promo_offers/dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type ActionDal interface {
	GetAllCount(ctx context.Context) (int64, error)
	GetCooperationTypeCounts(ctx context.Context) ([]dto.CountItem, error)
}

type CommonDal interface {
	GetCooperationTypesByIDs(ctx context.Context, ids []int64) (map[int64]string, error)
}

type Action struct {
	repository ActionDal
	commonDal  CommonDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		repository: actionDal.NewRepository(pool),
		commonDal:  commonDal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) (*dto.Response, error) {
	logger.Message(ctx, "[PromoOffers][FilterSettings] Получение доступных фильтров")

	allCount, err := a.repository.GetAllCount(ctx)
	if err != nil {
		return nil, err
	}

	cooperationTypeCounts, err := a.repository.GetCooperationTypeCounts(ctx)
	if err != nil {
		return nil, err
	}

	ids := lo.Map(cooperationTypeCounts, func(item dto.CountItem, _ int) int64 {
		return item.ID
	})
	cooperationTypesMap, err := a.commonDal.GetCooperationTypesByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	resp := &dto.Response{
		All:              allCount,
		CooperationTypes: make([]dto.CountItem, 0, len(cooperationTypeCounts)),
	}

	filt := lo.FilterMap(cooperationTypeCounts, func(item dto.CountItem, _ int) (dto.CountItem, bool) {
		name, ok := cooperationTypesMap[item.ID]
		if !ok {
			return dto.CountItem{}, false
		}

		item.Name = name
		return item, true
	})
	resp.CooperationTypes = append(resp.CooperationTypes, filt...)

	return resp, nil
}
