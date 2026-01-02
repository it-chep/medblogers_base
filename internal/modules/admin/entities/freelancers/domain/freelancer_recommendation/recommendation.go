package freelancer_recommendation

import "github.com/samber/lo"

type FreelancerRecommendation struct {
	id           int64
	freelancerID int64
	doctorID     int64
}

type FreelancerRecommendations []FreelancerRecommendation

func (r FreelancerRecommendations) DoctorIDs() []int64 {
	return lo.Map(r, func(item FreelancerRecommendation, _ int) int64 {
		return item.doctorID
	})
}

func NewRecommendation(id, doctorID, freelancerID int64) FreelancerRecommendation {
	return FreelancerRecommendation{
		id:           id,
		freelancerID: freelancerID,
		doctorID:     doctorID,
	}
}
