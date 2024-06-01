package main

import (
	"github.com/Ignite-Laboratories/JanOS/common"
	"github.com/Ignite-Laboratories/JanOS/common/config"
	"github.com/Ignite-Laboratories/JanOS/common/generate"
	"net/http"
	_ "net/http/pprof"
)

const spawnPoolSize = 2
const maxPerceptiveWidth = 32

func main() {
	go Profile()

	config.Initialize()
	o := common.NewObserver()
	go GenerateHeartbeat(o)
	go CreateNeuron(o)

	common.KeepAlive()
}

func Profile() {
	http.ListenAndServe("localhost:420", nil)
}

func CreateNeuron(o *common.Observer) {
	mm := common.NewStdMatchMaker(spawnPoolSize, maxPerceptiveWidth)
	o.ForwardTo(mm.InputStream)
}

func GenerateHeartbeat(o *common.Observer) {
	ng := generate.Noise(generate.NewNoiseType())
	ng.Broadcast()

	pg1 := generate.NewPulseGenerator()
	pg1.Pulse()

	pg2 := generate.NewPulseGenerator(generate.NewSeededPulseData())
	pg2.Pulse()

	pg3 := generate.NewPulseGenerator(generate.NewSeededPulseData())
	pg3.Pulse()

	go o.ReceiveFrom(ng.Output)
	go o.ReceiveFrom(pg1.Output)
	go o.ReceiveFrom(pg2.Output)
	go o.ReceiveFrom(pg3.Output)
}
