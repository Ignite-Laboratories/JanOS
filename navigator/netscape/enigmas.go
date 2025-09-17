package netscape

import (
	_ "embed"
	"net/http"
	"regexp"
	"strconv"

	"git.ignitelabs.net/core/sys/deploy"
	"git.ignitelabs.net/core/sys/log"
)

var Enigmas _enigmas

type _enigmas byte

func (_enigmas) Deploy() {
	deploy.Fly.Spark("exsx-enigmaneering-net", "navigator", "enigmas")
}

var subdomainRegex = regexp.MustCompile(`[A-Za-z0-9](?:[A-Za-z0-9\-]{0,61}[A-Za-z0-9])?`)
var enigmaRegex = regexp.MustCompile(`(?i)\be(\d+)(?:s(\d+))?\b`)

// Navigate executes a web server that listens on port 8080 (unless otherwise specified).
func (_enigmas) Navigate(repo string, port ...uint) {
	p := "4242"
	if len(port) > 0 {
		p = strconv.Itoa(int(port[0]))
	}

	if len(repo) > 0 && repo[len(repo)-1] == '/' {
		repo = repo[:len(repo)-1]
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		subdomain := subdomainRegex.FindString(r.Host)
		parts := enigmaRegex.FindStringSubmatch(subdomain)

		if len(parts) > 1 {
			redirect := repo + "/tree/main/enigma" + parts[1]

			if len(parts) > 2 && len(parts[2]) > 0 {
				redirect += "/solution" + parts[2]
			}
			log.Verbosef("exsx", "Navigating to '%s'\n", redirect)

			http.Redirect(w, r, redirect, http.StatusFound)
		}
	})

	addr := ":" + p
	log.Printf(ModuleName, "sparked *.enigmaneering.net%s\n", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf(ModuleName, err.Error())
	}
}
