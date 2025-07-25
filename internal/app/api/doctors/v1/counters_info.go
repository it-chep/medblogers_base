package v1

import (
	"encoding/json"
	"medblogers_base/internal/app/api/doctors/v1/dto/counters_info"
	"net/http"
)

// CountersInfo - /api/v1/counters_info [GET]
func (s *Service) CountersInfo(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из модуля
	countersDomain, err := s.doctors.Actions.CounterInfo.Do(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(counters_info.CountersResponse{
		DoctorsCount:     countersDomain.DoctorsCount,
		SubscribersCount: countersDomain.SubscribersCount,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
