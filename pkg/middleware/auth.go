package middleware

import (
	"net/http"
	"strings"

	"github.com/aweigor/nts-feed-generator/config"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, conf *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-API-TOKEN")
		diff := strings.Compare(token, conf.Auth.Secret)
		if diff != 0 {
			writeUnauthed(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
