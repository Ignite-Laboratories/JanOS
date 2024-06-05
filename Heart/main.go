package main

import (
	"Common/RPC/PerceptionAPI"
	"github.com/google/uuid"
	"time"
)

var ThisComponent = PerceptionAPI.NewComponent(uuid.New().String(), "tcp", "localhost:421")

func main() {
	c1 := ThisComponent.ConnectRemote("tcp", "localhost:420")
	c2 := ThisComponent.ConnectRemote("tcp", "localhost:420")
	c3 := ThisComponent.ConnectRemote("tcp", "localhost:420")

	go ProduceData(c1)
	go ProduceData(c2)
	go ProduceData(c3)
	select {}
}

func ProduceData(c *PerceptionAPI.Client) {
	for {
		c.ProcessPacket(c.ID + " Packet")
		time.Sleep(time.Second / 2)
	}
}
