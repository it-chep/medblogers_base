package v1

import (
	"context"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"

	"github.com/samber/lo"
)

// GetSMainSpecialities - /api/v1/main_specialities_list [GET]
func (i *Implementation) GetSMainSpecialities(ctx context.Context, _ *desc.GetMainSpecialitiesRequest) (*desc.MainSpecialitiesResponse, error) {
	specialities, err := i.doctors.Actions.MainSpecialitiesList.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.MainSpecialitiesResponse{
		Specialities: lo.Map(specialities, func(item *speciality.Speciality, _ int) *desc.MainSpecialitiesResponse_SpecialityItem {
			return &desc.MainSpecialitiesResponse_SpecialityItem{
				SpecialityId:   int64(item.ID()),
				SpecialityName: item.Name(),
			}
		}),
	}, nil
}
