package netscape

import (
	"context"
	_ "embed"
	"io/fs"
	"net"
	"net/http"
	"path"
	"strings"
	"sync"

	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/deploy"
	"git.ignitelabs.net/janos/core/sys/rec"
)

var Neural _neural

type _neural struct {
	moduleName string
	running    bool
	lock       sync.Mutex
}

func (*_neural) Deploy() {
	deploy.Fly.Spark("ignitelabs-net", "navigator", "ignite")
}

func (i *_neural) NavigateImpulse(imp *std.Impulse) {
	if i.running {
		return
	}

	i.lock.Lock()
	defer i.lock.Unlock()
	i.running = true
	i.moduleName = imp.Bridge

	port := "4242"
	addr := ":" + port
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	// Serve the embedded directory under /
	sub, err := fs.Sub(static, "ignite-src")
	if err != nil {
		rec.Fatalf(i.moduleName, err.Error())
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

	srv := &http.Server{
		Handler: handler,
	}

	imp.Cortex.Deferrals() <- func(wg *sync.WaitGroup) {
		_ = srv.Shutdown(context.Background())
		wg.Done()
	}

	go func() {
		if len(i.moduleName) == 0 {
			i.moduleName = ModuleName
		}

		http.Handle("/", handler)

		rec.Printf(i.moduleName, "sparked ignitelabs.net%s\n", addr)
		err = srv.Serve(ln)
		if err != nil {
			rec.Printf(i.moduleName, err.Error()+"\n")
		}
	}()
}
