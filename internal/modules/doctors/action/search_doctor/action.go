package search_doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/search_doctor/service/doctor"
)

// Action поиск доктора по фио, специальности, города
type Action struct {
	doctor *doctor.Service
}

// New .
func New() *Action {
	return &Action{
		doctor: doctor.NewSearchService(),
	}
}

func (a *Action) SearchDoctors(ctx context.Context, query string) error {
	return a.doctor.Search(ctx, query)
}
