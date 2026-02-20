package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	"medblogers_base/internal/pkg/logger"
)

// CheatersCount - /api/v1/blacklist_count [GET]
func (i *Implementation) CheatersCount(ctx context.Context, req *desc.CheatersCountRequest) (*desc.CheatersCountResponse, error) {
	count, err := i.doctors.Actions.BlackListCount.Do(ctx)
	if err != nil {
		logger.Error(ctx, "ошибка на проверке в черном списке", err)
		return &desc.CheatersCountResponse{CheatersCount: 0}, nil
	}

	return &desc.CheatersCountResponse{CheatersCount: count}, nil
}
