package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/id"
)

func main() {
	test(4)
	id.Next()
}

func test(a num.Nibble) {
	fmt.Println("here")
}
