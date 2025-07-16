package v1

import (
	"medblogers_base/internal/app/middleware"
	"medblogers_base/internal/modules/doctors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	doctors *doctors.Module
	router  *chi.Mux
}

func NewService(doctors *doctors.Module) *Service {
	s := &Service{
		doctors: doctors,
		router:  chi.NewRouter(),
	}
	s.setupMiddlewares()
	s.setupRoutes()

	return s
}

func (s *Service) setupRoutes() {
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/settings", s.Settings) // GET /api/v1/settings

		r.Route("/doctors", func(r chi.Router) {
			r.Get("/{doctor_id}", s.DoctorDetail)  // Обрабатывает /api/v1/doctors/23
			r.Get("/{doctor_id}/", s.DoctorDetail) // Обрабатывает /api/v1/doctors/23/
		})
	})
}

func (s *Service) setupMiddlewares() {
	s.router.Use(middleware.LoggerMiddleware)
}

func (s *Service) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
