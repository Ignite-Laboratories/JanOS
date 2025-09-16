package netscape

import (
	_ "embed"
	"log"
	"net/http"
	"strconv"
)

var Enigmas _enigmas

type _enigmas byte

// Navigate executes a web server that listens on port 8080 (unless otherwise specified).
func (_enigmas) Navigate(remote string, port ...uint) {
	p := "8080"
	if len(port) > 0 {
		p = strconv.Itoa(int(port[0]))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.)
		//repo := remote + r.URL.Path
		//
		//// Browser: redirect to GitHub. Use tree/HEAD for directories; it’s fine for blobs too.
		//http.Redirect(w, r, repo, http.StatusFound)
	})

	addr := ":" + p
	log.Printf("'*.enigmaneering.net' listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
