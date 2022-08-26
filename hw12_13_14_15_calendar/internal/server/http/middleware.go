package internalhttp

import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			start := time.Now()
			s.log.Info(fmt.Sprintf("%s %s %s %s %d %v %s", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, 200, time.Since(start), r.UserAgent())) // nolint: lll
		}()
		next.ServeHTTP(w, r)
	})
}
