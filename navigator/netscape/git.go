package netscape

import (
	"html/template"
	"net/http"
	"strconv"

	"git.ignitelabs.net/core/sys/deploy"
	"git.ignitelabs.net/core/sys/log"
)

// GitVanity drives the git.ignitelabs.net service, which acts as a "vanity URL" for Go imports.
//
// NOTE: The address 'git.ignitelabs.net' is implicit through the request and not present in the actual code.
//
// It's quite simple - just tell it where you want your hostname to redirect to and what port to listen on (default 8080)
//
// To facilitate vanity requests, you have to do two things:
//
// 0. Route your address (i.e. https://git.ignitelabs.net) to the Git repo (i.e. https://github.com/ignite-laboratories)
//
// 1. Handle the "go-get=1" query parameter and return the below template:
//
//	<html>
//	  <head>
//	    <meta name="go-import" content="[importPath] git [remote]">
//	    <meta name="go-source" content="[importPath] [remote] [remote]/tree/HEAD{/dir} [remote]/blob/HEAD{/dir}/{file}#L{line}">
//	  </head>
//	  <body>OK</body>
//	</html>
//
// That's really it!  No fancy libraries are needed, just a simple HTTP handler =)
var GitVanity _gitVanity

type _gitVanity byte

func (_gitVanity) Deploy() {
	deploy.Fly.Spark("git-ignitelabs-net", "navigator", "git")
}

// Navigate executes a web server that listens on port 4242 (unless otherwise specified).
func (_gitVanity) Navigate(remote string, port ...uint) {
	p := "4242"
	if len(port) > 0 {
		p = strconv.Itoa(int(port[0]))
	}

	metaTmpl := template.Must(template.New("meta").Parse(`<!doctype html>
<html><head>
<meta name="go-import" content="{{.ImportPath}} git {{.Remote}}">
<meta name="go-source" content="{{.ImportPath}} {{.Remote}} {{.Remote}}/tree/HEAD{/dir} {{.Remote}}/blob/HEAD{/dir}/{file}#L{line}">
</head><body>OK</body></html>`))

	type metaData struct {
		ImportPath string // e.g. git.ignitelabs.net/enigmaneering/enigma0
		Remote     string // e.g. https://github.com/ignite-laboratories/enigmaneering/enigma0
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo := remote + r.URL.Path

		// Go toolchain probe: serve meta tags (no redirect).
		if r.URL.Query().Get("go-get") == "1" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if err := metaTmpl.Execute(w, metaData{
				ImportPath: r.Host + r.URL.Path,
				Remote:     repo,
			}); err != nil {
				http.Error(w, "template error", http.StatusInternalServerError)
			}
			return
		}

		// Browser: redirect to GitHub. Use tree/HEAD for directories; it’s fine for blobs too.
		http.Redirect(w, r, repo, http.StatusFound)
	})

	addr := ":" + p
	log.Printf(ModuleName, "sparked git.ignitelabs.net%s\n", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf(ModuleName, err.Error())
	}
}
