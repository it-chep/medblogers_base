package spec

import (
	"context"
	"fmt"
)

// Specification представление правил бизнес-логики в виде цепочки объектов,
// связанных операциями булевой логики.
type Specification[T any] interface {
	IsSatisfied(ctx context.Context, t T) (bool, error)
	And(s Specification[T]) Specification[T]
	AndNot(s Specification[T]) Specification[T]
	Or(s Specification[T]) Specification[T]
	OrNot(s Specification[T]) Specification[T]
}

type specification[T any] struct {
	isSatisfied func(ctx context.Context, t T) (bool, error)
}

// IsSatisfied проверяет, соответствует ли объект спецификации
func (spec specification[T]) IsSatisfied(ctx context.Context, t T) (bool, error) {
	return spec.isSatisfied(ctx, t)
}

// And создает новую спецификацию путем объединения спецификаций с логической операцией AND
func (spec specification[T]) And(specification Specification[T]) Specification[T] {
	return NewSpecification[T](func(ctx context.Context, t T) (bool, error) {
		leftOk, leftErr := spec.isSatisfied(ctx, t)
		if !leftOk {
			return false, leftErr
		}

		rightOk, rightErr := specification.IsSatisfied(ctx, t)
		if !rightOk {
			return false, rightErr
		}

		return true, nil
	})
}

// AndNot создает новую спецификацию путем объединения спецификаций с логической операцией AND, переданная спецификация будет взята с отрицанием
func (spec specification[T]) AndNot(specification Specification[T]) Specification[T] {
	return NewSpecification[T](func(ctx context.Context, t T) (bool, error) {
		leftOk, leftErr := spec.isSatisfied(ctx, t)
		if !leftOk {
			return false, leftErr
		}

		rightOk, rightErr := specification.IsSatisfied(ctx, t)
		if rightOk || rightErr == nil {
			return false, fmt.Errorf("правило валидации, зарегистрированное с логическим отрицанием отработало без ошибки")
		}

		return true, nil
	})
}

// Or создает новую спецификацию путем объединения спецификаций с логической операцией OR
func (spec specification[T]) Or(specification Specification[T]) Specification[T] {
	return NewSpecification[T](func(ctx context.Context, t T) (bool, error) {
		leftOk, leftErr := spec.isSatisfied(ctx, t)
		if leftOk {
			return true, nil
		}

		rightOk, rightErr := specification.IsSatisfied(ctx, t)
		if rightOk {
			return true, nil
		}

		return false, newErrRightLeft(leftErr, rightErr)
	})
}

// OrNot создает новую спецификацию путем объединения спецификаций с логической операцией OR, переданная спецификация будет взята с отрицанием
func (spec specification[T]) OrNot(specification Specification[T]) Specification[T] {
	return NewSpecification[T](func(ctx context.Context, t T) (bool, error) {
		leftOk, leftErr := spec.isSatisfied(ctx, t)
		if leftOk {
			return true, nil
		}

		rightOk, rightErr := specification.IsSatisfied(ctx, t)
		if rightOk || rightErr == nil {
			return false, newErrRightLeft(leftErr, fmt.Errorf("правило валидации, зарегистрированное с логическим отрицанием отработало без ошибки"))
		}

		return true, nil
	})
}

// Not создает новую спецификацию путем меверсирования текущей
func Not[T any](specification Specification[T]) Specification[T] {
	return NewSpecification[T](func(ctx context.Context, t T) (bool, error) {
		ok, err := specification.IsSatisfied(ctx, t)
		if ok || err == nil {
			return false, fmt.Errorf("правило валидации, зарегистрированное с логическим отрицанием отработало без ошибки")
		}

		return true, nil
	})
}

// NewSpecification создает новую спецификацию
func NewSpecification[T any](isSatisfied func(ctx context.Context, t T) (bool, error)) Specification[T] {
	return specification[T]{isSatisfied: func(ctx context.Context, t T) (bool, error) {
		valid, err := isSatisfied(ctx, t)
		if err != nil {
			return valid, err
		}

		if !valid {
			return valid, fmt.Errorf("отсутсвует ошибка для правила валидации")
		}

		return true, nil
	}}
}
