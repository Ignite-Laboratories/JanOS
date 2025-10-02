package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/std/neural"
	"git.ignitelabs.net/janos/core/sys/deploy"
	"git.ignitelabs.net/janos/core/sys/rec"
)

//go:embed src/*
var static embed.FS

var port = "4242"
var cortex = std.NewCortex(std.RandomName())

func main() {
	if len(os.Args) > 1 && os.Args[1] == "deploy" {
		deploy.Fly.Spark("ignitelabs-net", "navigator", "ignite")
	} else {
		cortex.Frequency = 1 //hz
		cortex.Mute()

		cortex.Synapses() <- neural.Net.Server(life.Looping, "ignitelabs.net", ":4242", Handler, func(imp *std.Impulse) {
			cortex.Impulse()
		})

		cortex.Spark()
		cortex.Impulse()
		core.KeepAlive()
	}
}

func Handler(imp *std.Impulse) http.Handler {
	sub, err := fs.Sub(static, "src")
	if err != nil {
		rec.Fatalf(imp.Bridge.String(), err.Error())
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
