package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// DoctorDetail - /api/v1/doctors/{doctor_id} [GET]
func (s *Service) DoctorDetail(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из URL
	idStr := chi.URLParam(r, "doctor_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}

	// Получаем данные из модуля
	_, err = s.doctors.Actions.DoctorDetail.Do(r.Context(), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
}
