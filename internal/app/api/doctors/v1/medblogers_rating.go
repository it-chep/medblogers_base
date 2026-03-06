package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/doctors/action/medblogers_rating/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
)

func (i *Implementation) MedblogersRating(ctx context.Context, _ *desc.MedblogersRatingRequest) (*desc.MedblogersRatingResponse, error) {
	items, err := i.doctors.Actions.MedblogersRating.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.MedblogersRatingResponse{
		Doctors: lo.Map(items, func(item dto.RatingItem, index int) *desc.MedblogersRatingResponse_Doctor {
			return &desc.MedblogersRatingResponse_Doctor{
				Slug:     item.Slug,
				Name:     item.Name,
				MbcCoins: item.MBCCoins,
				City: &desc.CityItem{
					Id:   item.CityID,
					Name: item.CityName,
				},
				Speciality: &desc.SpecialityItem{
					Id:   item.SpecialityID,
					Name: item.SpecialityName,
				},
				Image: item.Image,
			}
		}),
	}, nil
}
