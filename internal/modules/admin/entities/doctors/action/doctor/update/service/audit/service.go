package audit

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
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
func (s *Service) Log(ctx context.Context, doc *doctor.Doctor) error {
	userID := pkgctx.GetUserIDFromContext(ctx)

	// Конвертим в json доктора
	json, err := doc.Json()
	if err != nil {
		return err
	}

	err = s.logger.Log(ctx, userID, audit_logger.LogData{
		Description: "Сохранение старых данных доктора",
		Body:        json,
		Action:      "Обновление доктора",
		EntityName:  audit_logger.DoctorEntity,
		EntityID:    int64(doc.GetID()),
	})
	if err != nil {
		return err
	}

	return nil
}
