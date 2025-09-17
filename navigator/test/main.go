package main

import (
	"os"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/navigator/netscape"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "deploy" {
		netscape.IgniteLabs.Deploy()
		netscape.GitVanity.Deploy()
		netscape.Enigmas.Deploy()
	} else {
		go netscape.IgniteLabs.Navigate(4242)
		go netscape.GitVanity.Navigate("https://GitHub.com/ignite-laboratories", 4243)
		go netscape.Enigmas.Navigate(4244)
	}

	core.WhileAlive()
}
