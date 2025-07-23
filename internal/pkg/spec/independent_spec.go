package spec

import (
	"context"
)

type rule[T any] func(ctx context.Context, t T) (bool, SpecError)

// IndependentSpecification независимая спецификация, в которой правила не зависят друг от друга. Собирает все ошибки, а не выходит при первой
type IndependentSpecification[T any] interface {
	Validate(ctx context.Context, t T)
	And(specification rule[T])

	Errors() []SpecError
}

type independentSpec[T any] struct {
	rules      []rule[T]
	specErrors []SpecError
}

// SpecError представление ошибки спецификации
type SpecError struct {
	Code    int64
	Message string
	Field   string
}

// AddError добавляет ошибку к остальным
func (spec independentSpec[T]) AddError(code int64, message, field string) {
	if spec.specErrors == nil {
		spec.specErrors = []SpecError{}
	}

	spec.specErrors = append(spec.specErrors, SpecError{
		Code:    code,
		Message: message,
		Field:   field,
	})
}

// Errors возвращает ошибки валидации
func (spec independentSpec[T]) Errors() []SpecError {
	return spec.specErrors
}

// Validate запускает все правила
func (spec independentSpec[T]) Validate(ctx context.Context, t T) {
	for _, r := range spec.rules {
		ok, specError := r(ctx, t)
		if !ok {
			spec.AddError(specError.Code, specError.Message, specError.Field)
		}
	}
}

// And создает новую спецификацию путем объединения спецификаций с логической операцией AND
func (spec independentSpec[T]) And(specification rule[T]) {
	spec.rules = append(spec.rules, specification)
}

// NewIndependentSpecification создает новую спецификацию
func NewIndependentSpecification[T any]() IndependentSpecification[T] {
	return independentSpec[T]{
		rules: make([]rule[T], 0),
	}
}
