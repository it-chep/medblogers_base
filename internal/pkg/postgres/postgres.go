package postgres

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// SQLExecutor описывает основные операции
type SQLExecutor interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

// pool wrapper
type poolWrapper struct {
	pool *pgxpool.Pool
}

// PoolWrapper ...
type PoolWrapper interface {
	SQLExecutor
}

// NewPoolWrapper ...
func NewPoolWrapper(pool *pgxpool.Pool) PoolWrapper {
	return &poolWrapper{
		pool: pool,
	}
}

// Exec ...
func (w *poolWrapper) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return w.pool.Exec(ctx, sql, arguments...)
}

// Query ...
func (w *poolWrapper) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return w.pool.Query(ctx, sql, args...)
}

// QueryRow ...
func (w *poolWrapper) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return w.pool.QueryRow(ctx, sql, args...)
}
