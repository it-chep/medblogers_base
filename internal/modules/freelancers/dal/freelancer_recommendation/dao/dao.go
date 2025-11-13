package dao

import domain "medblogers_base/internal/modules/freelancers/domain/freelancer_recommendation"

type FreelancerRecommendationDAO struct {
	ID           int64 `db:"id"`
	FreelancerID int64 `db:"freelancer_id"`
	DoctorID     int64 `db:"doctor_id"`
}

type FreelancerRecommendationsDAO []FreelancerRecommendationDAO

func (r FreelancerRecommendationDAO) ToDomain() domain.FreelancerRecommendation {
	return domain.NewRecommendation(r.ID, r.DoctorID, r.FreelancerID)
}

func (r FreelancerRecommendationsDAO) ToDomain() domain.FreelancerRecommendations {
	recoms := make([]domain.FreelancerRecommendation, 0, len(r))
	for _, rec := range r {
		recoms = append(recoms, rec.ToDomain())
	}
	return recoms
}
