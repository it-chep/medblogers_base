package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	"medblogers_base/internal/pkg/logger"
)

func (i *Implementation) CheckCheating(ctx context.Context, req *desc.CheckCheatingRequest) (*desc.CheckCheatingResponse, error) {
	if len(req.GetTelegram()) == 0 {
		return nil, nil
	}

	isInBlackList, err := i.doctors.Actions.BlackListCheck.Do(ctx, req.GetTelegram())
	if err != nil {
		logger.Error(ctx, "ошибка на проверке в черном списке", err)
		return &desc.CheckCheatingResponse{IsInBlacklist: false}, nil
	}

	return &desc.CheckCheatingResponse{IsInBlacklist: isInBlackList}, nil
}
