package counters_info

type CountersResponse struct {
	DoctorsCount     int64  `json:"doctors_count"`
	SubscribersCount string `json:"subscribers_count"`
}
