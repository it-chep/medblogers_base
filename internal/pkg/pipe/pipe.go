package pipe

import (
	"context"
	"fmt"
	"medblogers_base/internal/pkg/logger"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// Func шаблон функции
type Func[T any] func(ctx context.Context, value T) (T, error)

// AnywayFunc шаблон функции
type AnywayFunc[T any] func(ctx context.Context, value T, err error)

// FilterFunc шаблон функции
type FilterFunc[T any] func(ctx context.Context, value T) (filteredCtx context.Context, filtered T, empty bool, err error)

// OverFunc шаблон функции
type OverFunc[T any] func(ctx context.Context, value T, pipe Pipe[T]) (T, error)

// Pipe набор функций, которые должны выполнится для некоторой сущности
type Pipe[T any] []Func[T]

// With начинается конфигурации последовательности для некоторого значения
func With[T any](with Func[T]) Pipe[T] {
	return Pipe[T]{with}
}

// Each добавляет к последовательности новое звено
func Each[S ~[]T, T any](with Func[T]) Func[S] {
	return func(ctx context.Context, value S) (_ S, _ error) {
		for i, v := range value {
			v, err := with(ctx, v)
			if err != nil {
				return value, err
			}
			value[i] = v
		}
		return value, nil
	}
}

// EachAsync добавляет к последовательности новое звено
func EachAsync[S ~[]T, T any](with Func[T], limiter int) Func[S] {
	return func(ctx context.Context, value S) (S, error) {
		eg, egCtx := errgroup.WithContext(ctx)

		if limiter != 0 {
			eg.SetLimit(limiter)
		}

		part := func(i int) func() error {
			return func() (err error) {
				value[i], err = with(egCtx, value[i])
				return
			}
		}

		for i := range value {
			eg.Go(part(i))
		}

		return value, eg.Wait()
	}
}

// errPipeBreak покинуть пайп без ошибок
var errPipeBreak = errors.New("pipe break")

// Break покинуть пайп без ошибок
func Break() error {
	return errPipeBreak
}

// If начинает последовательность, выполняя проверку переданного условия
func If[T any](cond func(context.Context, T) (bool, error), then Pipe[T], elseRes Pipe[T]) Pipe[T] {
	return Pipe[T]{func(ctx context.Context, value T) (T, error) {
		res, err := cond(ctx, value)
		if err != nil {
			return value, err
		}

		if res {
			return then.Run(ctx, value).Get()
		}
		return elseRes.Run(ctx, value).Get()
	}}
}

// Over начинается конфигурации последовательности для некоторого значения
func Over[T any](over OverFunc[T], p Pipe[T]) Pipe[T] {
	return Pipe[T]{func(ctx context.Context, value T) (T, error) {
		return over(ctx, value, p)
	}}
}

// OverFiltered выполняет пайп с отфильтрованным value. Необходим доступ к мутабельным данным через указатели
func OverFiltered[T any](filter FilterFunc[T], p Pipe[T]) Pipe[T] {
	return Pipe[T]{func(ctx context.Context, value T) (T, error) {
		filteredCtx, filteredValue, empty, err := filter(ctx, value)
		if err != nil || empty {
			return value, err
		}

		return value, p.Run(filteredCtx, filteredValue).Err()
	}}
}

// With добавляет к последовательности новое звено
func (pipe Pipe[T]) With(with Func[T]) Pipe[T] {
	return append(pipe, with)
}

// IfFunc ...
type IfFunc[T any] func(context.Context, T) (bool, error)

// If добавляет к последовательности новое звено, выполняя проверку переданного условия
func (pipe Pipe[T]) If(cond IfFunc[T], then Pipe[T], elseRes Pipe[T]) Pipe[T] {
	return append(pipe, func(ctx context.Context, value T) (T, error) {
		res, err := cond(ctx, value)
		if err != nil {
			return value, err
		}

		if res {
			return then.Run(ctx, value).Get()
		}
		return elseRes.Run(ctx, value).Get()
	})
}

// IfCond добавляет к последовательности новое звено, выполняя проверку переданного условия
func (pipe Pipe[T]) IfCond(cond bool, then Pipe[T], elseRes Pipe[T]) Pipe[T] {
	return append(pipe, func(ctx context.Context, value T) (T, error) {
		if cond {
			return then.Run(ctx, value).Get()
		}
		return elseRes.Run(ctx, value).Get()
	})
}

// Over добавляет к последовательности новое звено, обобщенное общей логикой
func (pipe Pipe[T]) Over(f OverFunc[T], p Pipe[T]) Pipe[T] {
	return append(pipe, func(ctx context.Context, value T) (T, error) {
		return f(ctx, value, p)
	})
}

// OverFiltered выполняет пайп с отфильтрованным value. Необходим доступ к мутабельным данным через указатели
func (pipe Pipe[T]) OverFiltered(filter FilterFunc[T], p Pipe[T]) Pipe[T] {
	return append(pipe, func(ctx context.Context, value T) (T, error) {
		filteredCtx, filteredValue, empty, err := filter(ctx, value)
		if err != nil || empty {
			return value, err
		}

		return value, p.Run(filteredCtx, filteredValue).Err()
	})
}

func executeAnyways[T any](ctx context.Context, value T, err error, anyway ...AnywayFunc[T]) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	wg := sync.WaitGroup{}
	for _, f := range anyway {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f(ctx, value, err)
		}()
	}

	go func() {
		wg.Wait()
		cancel()
	}()
}

// Run запускает выполнение всех последовательностей возвращая финальный результат
func (pipe Pipe[T]) Run(ctx context.Context, value T) Observer[T] {
	var err error
	for _, with := range pipe {
		value, err = with(ctx, value)
		if errors.Is(err, errPipeBreak) {
			err = nil
			break
		}
		if err != nil {
			return Observer[T]{Result[T]{
				res: value,
				err: err,
			}}
		}
	}

	return Observer[T]{Result[T]{
		res: value,
		err: err,
	}}
}

// Result результат выполнения пайплайна
type Result[T any] struct {
	res T
	err error
}

// Get получение ответа
func (r Result[T]) Get() (T, error) {
	return r.res, r.err
}

// Must гарантированное получение ответа
func (r Result[T]) Must() T {
	return r.res
}

// Err ошибка выполнения
func (r Result[T]) Err() error {
	return r.err
}

// Observer ответ от пайпа
type Observer[T any] struct {
	Result[T]
}

// Anyway функция обрабатывающая итоговый ответ без блокировок
func (p Observer[T]) Anyway(ctx context.Context, action ...AnywayFunc[T]) Observer[T] {
	val, err := p.Get()
	logger.Message(ctx, fmt.Sprintf("[Anyway] активация вызова '%d' функций для обработки ошибки '%v'", len(action), err))
	executeAnyways(ctx, val, err, action...)
	return p
}
