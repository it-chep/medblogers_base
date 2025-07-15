package search_doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/search_doctor/dal"
	"medblogers_base/internal/pkg/postgres"

	"medblogers_base/internal/modules/doctors/action/search_doctor/service/doctor"
)

// Action поиск доктора по фио, специальности, города
type Action struct {
	doctor *doctor.Service
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		doctor: doctor.NewSearchService(dal.NewRepository(pool)),
	}
}

func (a *Action) SearchDoctors(ctx context.Context, query string) error {
	return a.doctor.Search(ctx, query)
}
