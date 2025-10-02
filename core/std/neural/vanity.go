package neural

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/std"
)

// Vanity drives the git.ignitelabs.net service, which acts as a "vanity URL" for Go imports.
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
func (_net) Vanity(lifecycle life.Cycle, source string, remote string, port uint, onDisconnect ...func(*std.Impulse)) std.Synapse {
	p := strconv.Itoa(int(port))

	metaTmpl := template.Must(template.New("meta").Parse(`<!doctype html>
<html><head>
<meta name="go-import" content="{{.ImportPath}} git {{.Remote}}.git">
<meta name="go-source" content="{{.ImportPath}} {{.Remote}} {{.Remote}}/tree/HEAD{/dir} {{.Remote}}/blob/HEAD{/dir}/{file}#L{line}">
</head><body>OK</body></html>`))

	type metaData struct {
		ImportPath string
		Remote     string
	}

	return Net.Server(lifecycle, source+":"+p, ":"+p, func(imp *std.Impulse) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			first := r.URL.Path
			after := ""
			if len(first) > 0 {
				if first[0] == '/' {
					first = first[1:]
				}
			}
			if i := strings.IndexByte(first, '/'); i >= 0 {
				f := first[:i]
				after = first[i:]
				first = f
			}
			repo := r.Host + "/" + first
			rem := remote + "/" + first

			if rem[len(rem)-1] == '/' {
				rem = rem[:len(rem)-1]
			}

			if r.URL.Query().Get("go-get") == "1" {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				if err := metaTmpl.Execute(w, metaData{
					ImportPath: repo,
					Remote:     rem,
				}); err != nil {
					http.Error(w, "template error", http.StatusInternalServerError)
				}
				return
			}

			if len(after) > 0 {
				http.Redirect(w, r, rem+"/tree/HEAD"+after, http.StatusFound)
			} else {
				http.Redirect(w, r, rem, http.StatusFound)
			}
		})
	}, onDisconnect...)
}
