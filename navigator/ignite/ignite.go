package ignite

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/rec"
)

//go:embed ignite-src/*
var static embed.FS

func Handler(imp *std.Impulse) http.Handler {
	sub, err := fs.Sub(static, "ignite-src")
	if err != nil {
		rec.Fatalf(imp.Bridge, err.Error())
	}
	fileServer := http.FileServer(http.FS(sub))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to serve the file if it exists
		upath := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		if upath == "" {
			upath = "index.html"
		}
		if _, err = fs.Stat(sub, upath); err == nil {
			fileServer.ServeHTTP(w, r)
			return
		}
		// Fallback to index.html for non-file paths
		if strings.Contains(upath, ".") {
			// Looks like an asset that genuinely doesn't exist
			http.NotFound(w, r)
			return
		}
		r2 := new(http.Request)
		*r2 = *r
		r2.URL.Path = "/index.html"
		fileServer.ServeHTTP(w, r2)
	})
}
