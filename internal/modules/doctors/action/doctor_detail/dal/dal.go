package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"medblogers_base/i
	"medblogers_base/internal/modules/doctors/dal/doctor_dal/dao"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Repository struct {
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository() *Repository {
	return &Repository{}
}

// GetDoctorInfo получает информацию о докторе
func (r Repository) GetDoctorInfo(ctx context.Context, doctorID int64) (*doctor.Doctor, error) {
	sql := fmt.Sprintf(`
				select id, name, slug, inst_url, vk_url, dzen_url from docstar_site_doctor where id = $1
				`)

	var doctorDAO dao.DoctorDAO
	if err := pgxscan.Select(ctx, r.db.Pool(ctx), &doctorDAO, sql, doctorID); err != nil {
		return nil, err
	}

	return doctorDAO.ToDomain(), nil
}
