package spec

import (
	"context"
)

type rule[T any] func(ctx context.Context, t T) (bool, error)

// IndependentSpecification независимая спецификация, в которой правила не зависят друг от друга. Собирает все ошибки, а не выходит при первой
type IndependentSpecification[T any] interface {
	Validate(ctx context.Context, t T) []error
	And(specification rule[T]) independentSpec[T]
}

type independentSpec[T any] struct {
	rules []rule[T]
}

// Validate запускает все правила
func (spec independentSpec[T]) Validate(ctx context.Context, t T) []error {
	errs := make([]error, 0)
	for _, r := range spec.rules {
		ok, specError := r(ctx, t)
		if !ok {
			errs = append(errs, specError)
		}
	}

	return errs
}

// And создает новую спецификацию путем объединения спецификаций с логической операцией AND
func (spec independentSpec[T]) And(specification rule[T]) independentSpec[T] {
	spec.rules = append(spec.rules, specification)
	return spec
}

// NewIndependentSpecification создает новую спецификацию
func NewIndependentSpecification[T any]() IndependentSpecification[T] {
	return independentSpec[T]{
		rules: make([]rule[T], 0),
	}
}
