package audit_logger

import (
	"context"
	"medblogers_base/internal/pkg/postgres"
)

type EntityName string

const (
	DoctorEntity     EntityName = "doctor"
	FreelancerEntity EntityName = "freelancer"
	DoctorVipEntity  EntityName = "doctor_vip"
)

type LogData struct {
	Description string
	Body        []byte // string ?? ?? ??
	Action      string
	EntityName  EntityName
	EntityID    int64
}

type Logger struct {
	pool postgres.PoolWrapper
}

// NewLogger создает новый аудит логгер
func NewLogger(db postgres.PoolWrapper) *Logger {
	return &Logger{
		pool: db,
	}
}

func (l *Logger) Log(ctx context.Context, userID int64, data LogData) error {
	sql := `
		insert into admin_audit (user_id, description, body, action, entity_name, entity_id) 
		values ($1, $2, $3, $4, $5, $6)
	`

	args := []interface{}{
		userID,
		data.Description,
		data.Body,
		data.Action,
		data.EntityName,
		data.EntityID,
	}

	_, err := l.pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
