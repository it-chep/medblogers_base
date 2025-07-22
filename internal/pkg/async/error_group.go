package async

import (
	"context"
	"fmt"
	"medblogers_base/internal/pkg/logger"

	eg "golang.org/x/sync/errgroup"
)

type ErrGroup struct {
	g   *eg.Group
	ctx context.Context
}

func NewErrGroup() *ErrGroup {
	g, ctx := eg.WithContext(context.Background())
	return &ErrGroup{
		g:   g,
		ctx: ctx,
	}
}

func WithContext(ctx context.Context) (*ErrGroup, context.Context) {
	g, retCtx := eg.WithContext(ctx)
	return &ErrGroup{g: g, ctx: retCtx}, retCtx
}

func (g *ErrGroup) Go(f func() error) {
	g.goWithContext(func(_ context.Context) error {
		return f()
	})
}

func (g *ErrGroup) GoWithContext(f func(ctx context.Context) error) {
	g.goWithContext(func(ctx context.Context) error {
		return f(ctx)
	})
}

func (g *ErrGroup) Wait() error {
	return g.g.Wait()
}

func (g *ErrGroup) SetLimit(n int) {
	g.g.SetLimit(n)
}

func (g *ErrGroup) goWithContext(f func(ctx context.Context) error) {
	g.g.Go(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = g.error(g.ctx, r)
			}
		}()

		err = f(g.ctx)
		return
	})
}

func (g *ErrGroup) error(ctx context.Context, recoverResult any) error {
	err := fmt.Errorf("%v", recoverResult)

	logger.Error(ctx, "Ошибка в errGroup", err)
	return err
}
