package middleware

import (
	"net/http"
	"net/http/httputil"
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	log *zap.Logger
}

func NewLogger(log *zap.Logger) Logger {
	if log == nil {
		log = zap.NewNop()
	}

	return Logger{
		log: log.With(zap.String("go.component", "middleware.Logger")),
	}
}

func (l Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := l.log

		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Error("failed to dump incoming request", zap.Error(err))
		}

		log = log.With(zap.ByteString("debug.request.data", reqDump))

		debugRW := debugResponseWriter{ResponseWriter: w}
		start := time.Now()
		next.ServeHTTP(&debugRW, r)

		log.Debug("request served",
			zap.ByteString("debug.response.body", debugRW.body),
			zap.Int("debug.response.status_code", debugRW.status),
			zap.Duration("debug.request.duration", time.Since(start)),
		)
	})
}

type debugResponseWriter struct {
	http.ResponseWriter

	status int
	body   []byte
}

func (w *debugResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *debugResponseWriter) Write(body []byte) (int, error) {
	w.body = body
	return w.ResponseWriter.Write(body)
}
