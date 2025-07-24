package config

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type configKey string

const contextConfigKey configKey = "config"

type Config interface {
	GetValue(ctx context.Context, key string) (Value, error)
}

type config struct {
	pool *pgxpool.Pool
}

// New ...
func New(pool *pgxpool.Pool) Config {
	return &config{
		pool: pool,
	}
}

func (c *config) GetValue(ctx context.Context, key string) (Value, error) {
	sql := `select value from config where key = $1`

	var jsonData []byte
	err := c.pool.QueryRow(ctx, sql, key).Scan(&jsonData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get config value")
	}

	var val concreteValue
	if err = json.Unmarshal(jsonData, &val); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config value")
	}

	return &val, nil
}

// ContextWithConfig прокинуть конфиг в контекст
func ContextWithConfig(ctx context.Context, cfg Config) context.Context {
	return context.WithValue(ctx, contextConfigKey, cfg)
}

func fromContext(ctx context.Context) Config {
	l, ok := ctx.Value(contextConfigKey).(Config)
	if ok {
		return l
	}
	return nil
}

func GetValue(ctx context.Context, key string) (Value, error) {
	cfg := fromContext(ctx)
	return cfg.GetValue(ctx, key)
}
