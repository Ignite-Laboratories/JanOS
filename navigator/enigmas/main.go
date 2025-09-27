package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/std/neural"
	"git.ignitelabs.net/janos/core/sys/deploy"
	"git.ignitelabs.net/janos/core/sys/rec"
)

var port = "4242"
var cortex = std.NewCortex(std.RandomName())

func main() {
	if len(os.Args) > 1 && os.Args[1] == "deploy" {
		deploy.Fly.Spark("exsx-enigmaneering-net", "navigator", "enigmas")
	} else {
		cortex.Frequency = 1 //hz
		cortex.Mute()

		cortex.Synapses() <- neural.Net.Server("enigmaneering.net", ":4242", Handler)

		cortex.Spark()
		cortex.Impulse()
		core.KeepAlive()
	}
}

var subdomainRegex = regexp.MustCompile(`[A-Za-z0-9](?:[A-Za-z0-9\-]{0,61}[A-Za-z0-9])?`)
var enigmaRegex = regexp.MustCompile(`(?i)\be(\d+)(?:s(\d+))?\b`)

func Handler(imp *std.Impulse) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		subdomain := subdomainRegex.FindString(r.Host)
		parts := enigmaRegex.FindStringSubmatch(subdomain)

		if len(parts) > 1 {
			redirect := "https://github.com/ignite-laboratories/enigmaneering" + "/tree/main/enigma" + parts[1]

			if len(parts) > 2 && len(parts[2]) > 0 {
				redirect += "/solution" + parts[2]
			}
			rec.Verbosef("exsx", "Navigating to '%s'\n", redirect)

			http.Redirect(w, r, redirect, http.StatusFound)
		} else {
			_, _ = fmt.Fprintf(w, "Hello, world!")
		}
	})
}
