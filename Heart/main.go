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

	rstr := c.Echo("TEST")
	fmt.Println(rstr)

	robj := c.GetObject("ASDF")
	js, _ := json.Marshal(robj)
	fmt.Println(string(js))

	c.ReceiveObject(robj)
}
