package main

import (
	"os"

	"git.ignitelabs.net/navigator/netscape"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "deploy" {
		netscape.IgniteLabs.Deploy()
	} else {
		netscape.IgniteLabs.Navigate()
	}
}
