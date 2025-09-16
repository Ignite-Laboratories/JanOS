package pkg

import (
	"fmt"

	"git.ignitelabs.net/core"
)

func GetDir() {
	fmt.Println(core.RelativePath("navigator", "ignite"))
}
