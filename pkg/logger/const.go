package logger

import "context"

type ctxKey int

const (
	requestIdKey ctxKey = iota
)

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIdKey, id)
}

func RequestIDFrom(ctx context.Context) string {
	if v, ok := ctx.Value(requestIdKey).(string); ok {
		return v
	}
	return ""
}
