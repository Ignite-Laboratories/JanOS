package main

import (
	"git.ignitelabs.net/core/sys/deploy"
)

func main() {
	deploy.Fly.Spark("git-ignitelabs-net", "navigator", "git")
	deploy.Fly.Spark("ignitelabs-net", "navigator", "ignite")
}
