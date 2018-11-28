package middleware

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
)

const loggerKey = "LOGGER"

type ResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func NewResponseWriter(w http.ResponseWriter) ResponseWriter  {
	return ResponseWriter{w, 200}
}

func Logger(ctx context.Context) *zap.Logger {
	if ctx.Value(loggerKey) == nil {
		panic(errors.New("logger not available"))
	}

	return ctx.Value(loggerKey).(*zap.Logger)
}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LoggerMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	logger, _ := zap.NewProduction()
	//defer logger.Sync()
	return func(w http.ResponseWriter, r *http.Request) {
		sugar := logger.Sugar()
		customRW := NewResponseWriter(w)
		handlerFunc(&customRW, r.WithContext(WithLogger(r.Context(), logger)))
		sugar.Infow("request",
			"status", customRW.status,
			"method", r.Method,
			"url", r.URL.Path,
		)
	}
}

