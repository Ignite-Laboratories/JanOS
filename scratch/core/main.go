package main

import "git.ignitelabs.net/core"

func main() {
	core.Synapses <- core.Spark{}
	core.Ignite()
}
