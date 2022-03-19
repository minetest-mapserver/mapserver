package web

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func CachedServeFunc(h http.HandlerFunc) http.HandlerFunc {
	var etag = fmt.Sprintf(`"%d"`, time.Now().UnixMicro())
	return func(w http.ResponseWriter, r *http.Request) {
		if match := r.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		w.Header().Set("Cache-Control", "max-age=60")
		w.Header().Set("ETag", etag)
		h.ServeHTTP(w, r)
	}
}
