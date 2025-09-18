package main

import "git.ignitelabs.net/core"

func main() {
	core.NeuralActivity <- core.Spark{}
	core.Ignite()
}
