package main

import (
	"Common/RPC/PerceptionAPI"
	"time"
)

func main() {
	go AddClient("A")
	go AddClient("B")
	go AddClient("C")

	select {}
}

func AddClient(id string) {
	//c := PerceptionAPI.NewClient("unix", "/tmp/perception.sock")
	c := PerceptionAPI.NewClient("tcp", "localhost:420")
	go c.Start()

	time.Sleep(time.Second)
	for {
		c.ProcessPacket(id)
		time.Sleep(time.Second / 4)
	}
}
