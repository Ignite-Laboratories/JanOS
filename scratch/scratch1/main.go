package main

import (
	"fmt"
)

func main() {
	base := 256
	for i := 2; i < base; i++ {
		fmt.Println(fmt.Sprintf("%x", i))
	}
}
