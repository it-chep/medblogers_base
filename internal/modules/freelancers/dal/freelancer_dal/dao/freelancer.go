package dao

type FreelancerMiniature struct {
}

type PriceCategory struct {
	ID               int64 `db:"id" json:"id"`
	FreelancersCount int64 `db:"freelancers_count" json:"freelancers_count"`
}
