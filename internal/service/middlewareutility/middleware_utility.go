package middlewareutility

import (
	"fmt"
	"net/http"

	"go.uber.org/zap/zapcore"

	"github.com/ardihikaru/go-chi-example-part-1/pkg/logger"
)

type Service struct {
	log *logger.Logger
}

// NewService creates a new timeout handler service
func NewService(log *logger.Logger) *Service {
	return &Service{
		log: log,
	}
}

// Log logs message based on the log level
func (svc *Service) Log(level zapcore.Level, msg string) {
	switch level {
	case zapcore.DebugLevel:
		svc.log.Debug(msg)
	case zapcore.InfoLevel:
		svc.log.Info(msg)
	case zapcore.WarnLevel:
		svc.log.Warn(msg)
	case zapcore.ErrorLevel:
		svc.log.Error(msg)
	case zapcore.FatalLevel:
		svc.log.Fatal(msg)
	case zapcore.PanicLevel:
		svc.log.Panic(msg)
	default:
		svc.log.Error(fmt.Sprintf("unexpected log level type: %d", level))
	}
}

// AuthorizeAccess is a middleware to enforce function's access control of the requester
func (svc *Service) AuthorizeAccess(resourceCode string, act string) func(next http.Handler) http.Handler {
	//TODO implement me
	panic("implement me")
}
