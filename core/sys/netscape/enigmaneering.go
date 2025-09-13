package netscape

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var Enigmaneering _enigmaneering

type _enigmaneering byte

func (_enigmaneering) Navigate() {
	target, _ := url.Parse("https://github.com")

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Rewrite scheme/host
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host

			// Prepend the org segment: /ignite-laboratories
			// Keep the incoming path as-is after the leading slash.
			req.URL.Path = joinURLPath("/ignite-laboratories", req.URL.Path)

			// Optionally preserve X-Forwarded-* headers
			// (ReverseProxy will set X-Forwarded-For automatically)
			if req.Header.Get("X-Forwarded-Proto") == "" {
				req.Header.Set("X-Forwarded-Proto", "https")
			}
		},

		// Optional: if GitHub redirects, rewrite Location so clients still see your host
		// Remove this block if you don’t want that behavior.
		ModifyResponse: func(resp *http.Response) error {
			loc := resp.Header.Get("Location")

			// Example: rewrite back to your proxy host name if desired
			u, err := url.Parse(loc)
			u.Scheme = "https"
			u.Host = "git.ignitelabs.net"
			resp.Header.Set("Location", u.String())

			if loc == "" {
				return nil
			}
			if err != nil {
				return nil
			}
			return nil
		},
	}

	// Serve all requests through the proxy.
	// Typically you run this behind a server that receives Host: git.ignitelabs.net
	http.Handle("/", proxy)

	log.Println("proxy listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
