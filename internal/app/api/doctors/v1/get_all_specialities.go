package v1

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"

	"github.com/samber/lo"
)

// GetSpecialities - /api/v1/specialities_list [GET]
func (i *Implementation) GetSpecialities(ctx context.Context, _ *desc.GetSpecialitiesRequest) (*desc.SpecialitiesResponse, error) {
	specialities, err := i.doctors.Actions.AllSpecialities.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.SpecialitiesResponse{
		Specialities: lo.Map(specialities, func(item *speciality.Speciality, _ int) *desc.SpecialitiesResponse_SpecialityItem {
			return &desc.SpecialitiesResponse_SpecialityItem{
				SpecialityId:   int64(item.ID()),
				SpecialityName: item.Name(),
			}
		}),
	}, nil
}
