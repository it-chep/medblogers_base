package v1

import (
	"context"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"

	"github.com/samber/lo"
)

func (i *Implementation) GetSpecialities(ctx context.Context, _ *desc.GetSpecialitiesRequest) (*desc.SpecialitiesResponse, error) {
	specialities, err := i.freelancers.Actions.GetAllSpecialities.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.SpecialitiesResponse{
		Specialities: lo.Map(specialities, func(item *speciality.Speciality, _ int) *desc.SpecialitiesResponse_SpecialityItem {
			return &desc.SpecialitiesResponse_SpecialityItem{
				SpecialityId:   item.ID(),
				SpecialityName: item.Name(),
			}
		}),
	}, nil
}
