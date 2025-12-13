package users

import (
	"context"
	"medblogers_base/internal/modules/auth/dal/users/dao"
	"medblogers_base/internal/modules/auth/domain/user"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// GetUserByEmail .
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	sql := `select * from go_users where email = $1`

	var usr dao.UserDAO
	err := pgxscan.Get(ctx, r.db, &usr, sql, email)
	if err != nil {
		return nil, err
	}

	return usr.ToDomain(), nil
}

// IsEmailExists .
func (r *Repository) IsEmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, "select exists (select 1 from go_users where email=$1 and password is null)", email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// SavePass .
func (r *Repository) SavePass(ctx context.Context, email, password string) error {
	_, err := r.db.Exec(ctx, "update go_users set password=$1 where email=$2 and password is null", password, email)
	return err
}
