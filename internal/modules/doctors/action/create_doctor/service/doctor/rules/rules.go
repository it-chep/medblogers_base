package rules

import (
	"context"
	"errors"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/spec"
)

// todo
// RuleValidTgChannelURL проверяет валидность тгк
var RuleValidTgChannelURL = func() spec.Specification[*doctor.Doctor] {
	return spec.NewSpecification(func(_ context.Context, doc *doctor.Doctor) (bool, error) {
		if doc.GetTgChannelURL() {
			// todo доменная ошибка
			return false, errors.New("не валиден")
		}

		return true, nil
	})
}
