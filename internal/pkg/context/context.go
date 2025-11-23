package context

import (
	"context"
	"net/url"

	"google.golang.org/grpc/metadata"
)

// Key type for context key
type Key string

const (
	userEmailKey Key = "email"
)

// WithValues returns ctx with values
func WithValues(ctx context.Context, values url.Values) context.Context {
	if len(values) == 0 {
		return ctx
	}
	for k, v := range values {
		if len(v) == 0 {
			continue
		}
		ctx = context.WithValue(ctx, Key(k), v[0])
	}
	return ctx
}

// WithValuesStandard returns ctx with values
func WithValuesStandard(ctx context.Context, values map[string]string) context.Context {
	if len(values) == 0 {
		return ctx
	}
	for k, v := range values {
		ctx = context.WithValue(ctx, Key(k), v)
	}
	return ctx
}

// GetValuesFromContext returns values by keys
func GetValuesFromContext(ctx context.Context, keys ...Key) map[Key]string {
	values := make(map[Key]string, len(keys))
	for _, key := range keys {
		if value, ok := ctx.Value(key).(string); ok {
			values[key] = value
		}
	}
	return values
}

// WithMeta get value from incoming meta by keys and append it to outgoing meta
func WithMeta(ctx context.Context, values map[string]string) context.Context {
	for key, value := range values {
		ctx = metadata.AppendToOutgoingContext(ctx, key, value)
	}
	return ctx
}

func WithEmailContext(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, userEmailKey, email)
}

func GetEmailFromContext(ctx context.Context) string {
	email, _ := ctx.Value(userEmailKey).(string)
	return email
}
