package audit

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/audit_logger"
	pkgctx "medblogers_base/internal/pkg/context"
)

type Logger interface {
	Log(ctx context.Context, userID int64, data audit_logger.LogData) error
}

type Service struct {
	logger Logger
}

func New(logger Logger) *Service {
	return &Service{
		logger: logger,
	}
}

// Log запись аудит лога
func (s *Service) Log(ctx context.Context, frlncr *freelancer.Freelancer) error {
	userID := pkgctx.GetUserIDFromContext(ctx)

	// Конвертим в json доктора
	json, err := frlncr.Json()
	if err != nil {
		return err
	}

	err = s.logger.Log(ctx, userID, audit_logger.LogData{
		Description: "Сохранение старых данных фрилансера",
		Body:        json,
		Action:      "Обновление фрилансера",
		EntityName:  audit_logger.FreelancerEntity,
		EntityID:    frlncr.GetID(),
	})
	if err != nil {
		return err
	}

	return nil
}
