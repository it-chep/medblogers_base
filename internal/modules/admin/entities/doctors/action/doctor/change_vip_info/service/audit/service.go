package audit

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/vip_card"
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
func (s *Service) Log(ctx context.Context, vip *vip_card.VipCard) error {
	if vip == nil {
		return nil
	}

	userID := pkgctx.GetUserIDFromContext(ctx)

	// Конвертим в json доктора
	json, err := vip.Json()
	if err != nil {
		return err
	}

	err = s.logger.Log(ctx, userID, audit_logger.LogData{
		Description: "Сохранение старых данных ВИП карточки",
		Body:        json,
		Action:      "Обновление випки доктора",
		EntityName:  audit_logger.DoctorVipEntity,
		EntityID:    vip.GetDoctorID(),
	})
	if err != nil {
		return err
	}

	return nil
}
