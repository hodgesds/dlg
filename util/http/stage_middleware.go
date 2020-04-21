package util

import (
	"net/http"

	"github.com/hodgesds/dlg/config"
	"gopkg.in/yaml.v2"
)

const (
	// StageMiddlewareDebugKey is the debug key for stage HTTP middleware.
	StageMiddlewareDebugKey = "stage_config"
)

// StageMiddleware is HTTP middleware for generating stage configs. It works by
// if the HTTP parameter stage_config is set.
func StageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		k := r.URL.Query().Get(StageMiddlewareDebugKey)
		if k == "" {
			if next != nil {
				next.ServeHTTP(w, r)
			}
			return
		}
		stage, err := config.StageFrom(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b, err := yaml.Marshal(stage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "text/yaml")
	})
}
