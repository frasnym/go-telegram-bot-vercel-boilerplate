package ctxdata

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type key string

const correlationIDKey key = "correlationID"

func GetCorrelationID(ctx context.Context) string {
	id, ok := ctx.Value(correlationIDKey).(string)
	if !ok {
		return "invalid"
	}
	return id
}

func EnsureCorrelationIDExist(r *http.Request) context.Context {
	ctx := r.Context()
	correlationID := r.Header.Get("Correlation-ID")

	if correlationID != "" {
		return context.WithValue(ctx, correlationIDKey, correlationID)
	}

	if _, ok := ctx.Value(correlationIDKey).(string); !ok {
		// If the correlation ID doesn't exist, generate a new one.
		newID := generateCorrelationID()
		ctx = context.WithValue(ctx, correlationIDKey, newID)
	}
	return ctx
}

func generateCorrelationID() string {
	// Replace this with your logic for generating a unique correlation ID.
	return uuid.NewString()
}
