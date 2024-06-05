package main

import (
	"Common/RPC/SampleAPI"
	"encoding/json"
	"fmt"
)

func main() {
	c := SampleAPI.NewClient("localhost:420")
	c.Start()

	str := c.Echo("Rawrrr")
	fmt.Println(str)

	obj := c.GetObject("ASDF")
	js, _ := json.Marshal(obj)
	fmt.Println(string(js))

	c.ReceiveObject(obj)
}
