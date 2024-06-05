package main

import (
	"Common/RPC/TestAPI"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	c := TestAPI.NewClient("localhost:420")

	time.Sleep(time.Millisecond * 250)

	str := c.Echo("Rawrrr")
	fmt.Println(str)

	obj := c.GetObject("ASDF")
	js, _ := json.Marshal(obj)
	fmt.Println(string(js))

	c.ReceiveObject(obj)
}
