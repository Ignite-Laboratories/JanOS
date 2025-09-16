package main

import (
	"os"

	"git.ignitelabs.net/navigator/netscape"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "deploy" {
		netscape.GitVanity.Deploy()
	} else {
		netscape.GitVanity.Navigate("https://github.com/ignite-laboratories")
	}
}
