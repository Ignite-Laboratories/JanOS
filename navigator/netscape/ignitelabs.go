package netscape

import (
	"embed"
	_ "embed"
	"io/fs"
	"net/http"
	"path"
	"strconv"
	"strings"

	"git.ignitelabs.net/core/sys/deploy"
	"git.ignitelabs.net/core/sys/log"
)

//go:embed ignite-src/*
var static embed.FS

var IgniteLabs _igniteLabs

type _igniteLabs byte

func (_igniteLabs) Deploy() {
	deploy.Fly.Spark("ignitelabs-net", "navigator", "ignite")
}

// Navigate executes a web server that listens on port 4242 (unless otherwise specified).
func (_igniteLabs) Navigate(port ...uint) {
	p := "4242"
	if len(port) > 0 {
		p = strconv.Itoa(int(port[0]))
	}
	// Serve the embedded directory under /
	sub, err := fs.Sub(static, "ignite-src")
	if err != nil {
		log.Fatalf(ModuleName, err.Error())
	}
	fileServer := http.FileServer(http.FS(sub))

	// Optional SPA-style fallback to index.html for client-side routes
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	http.Handle("/", handler)

	addr := ":" + p
	log.Printf(ModuleName, "sparked ignitelabs.net%s\n", addr)
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf(ModuleName, err.Error())
	}
}
