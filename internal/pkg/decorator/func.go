package decorator

import (
	"context"
)

type Action func(ctx context.Context) error

type With func(ctx context.Context, next Action) error

// ExecuteWith executes function with passed decorators. Order is important
func ExecuteWith(ctx context.Context, fnc Action, withs ...With) error {
	return chain(withs...)(ctx, fnc)
}

func chain(withs ...With) With {
	n := len(withs)
	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, next Action) error {
			var (
				chainFunc Action
				curI      int
			)
			chainFunc = func(ctx context.Context) error {
				if curI == lastI {
					return next(ctx)
				}
				curI++
				err := withs[curI](ctx, chainFunc)
				curI--
				return err

			}
			return withs[0](ctx, chainFunc)
		}
	}

	if n == 1 {
		return withs[0]
	}

	return func(ctx context.Context, next Action) error {
		return next(ctx)
	}
}
