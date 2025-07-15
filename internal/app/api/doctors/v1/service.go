package v1

import (
	"github.com/go-chi/chi/v5"
	"medblogers_base/internal/modules/doctors"
	"net/http"
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

	s.setupRoutes()
	return s
}

func (s *Service) setupRoutes() {
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/settings", s.Settings) // GET /api/v1/settings
		r.Route("/{doctorID}", func(r chi.Router) {
			r.Get("/", s.DoctorDetail) // GET /api/v1/doctors/{id}
		})
	})
}

func (s *Service) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
