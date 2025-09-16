package main

import (
	"navigator/pkg"
)

var test = false

func main() {
	if test {
		pkg.GetDir()
	}
}
