package dal

import "context"

type Repository struct {
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) FilterDoctors(ctx context.Context) error {
	return nil
}
