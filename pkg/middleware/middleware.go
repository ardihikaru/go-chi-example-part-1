package middleware

import (
	"net/http"

	"go.uber.org/zap/zapcore"
)

const (
	RequestId = "X-Request-Id"
)

// utility provides the interface for the functionality of logger.Logger
type utility interface {
	AuthorizeAccess(resourceCode string, act string) func(next http.Handler) http.Handler
	Log(level zapcore.Level, msg string)
}

type Resource struct {
	utility utility
}

// NewMiddleware creates a new middleware
func NewMiddleware(utility utility) *Resource {
	return &Resource{
		utility: utility,
	}
}
