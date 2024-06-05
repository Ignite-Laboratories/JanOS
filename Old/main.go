package main

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/common"
	"github.com/Ignite-Laboratories/JanOS/common/config"
	"github.com/Ignite-Laboratories/JanOS/common/generate"
	"net/http"
	_ "net/http/pprof"
)

const spawnPoolSize = 64
const maxPerceptiveWidth = 32

func main() {
	go Profile()

	config.Initialize()
	o := common.NewObserver()
	go CreateNeuron(o)
	go GenerateHeartbeat(o)

	common.KeepAlive()
}

func Profile() {
	http.ListenAndServe("localhost:420", nil)
}

func CreateNeuron(o *common.Observer) common.MatchMaker {
	mm := common.NewStdMatchMaker(spawnPoolSize, maxPerceptiveWidth)
	go o.ForwardTo(mm.InputStream)
	return mm
}

func GenerateHeartbeat(o *common.Observer) {
	//ng := generate.NewNoiseGenerator(generate.NewNoiseType())
	//go ng.Broadcast()

	pg1 := generate.NewPulseGenerator()
	go pg1.SpreadPulse(10)
	//go pg1.Pulse()

	v2 := fmt.Sprintf("5b75a8b7-11f2-%s-%s-%s", common.RandomString(4), common.RandomString(4), common.RandomString(12))
	pd2 := generate.NewValuedPulseData(v2)
	pg2 := generate.NewPulseGenerator(pd2)
	go pg2.SpreadPulse(10)
	//go pg2.Pulse()

	v3 := fmt.Sprintf("5b75a8b7-11f2-%s-%s-%s", common.RandomString(4), common.RandomString(4), common.RandomString(12))
	pd3 := generate.NewValuedPulseData(v3)
	pg3 := generate.NewPulseGenerator(pd3)
	go pg3.SpreadPulse(10)
	//go pg3.Pulse()

	//go o.ReceiveFrom(ng.Output)
	go o.ReceiveFrom(pg1.Output)
	go o.ReceiveFrom(pg2.Output)
	go o.ReceiveFrom(pg3.Output)
}
