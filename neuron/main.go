package main

import (
	"github.com/Ignite-Laboratories/JanOS/common"
	"github.com/Ignite-Laboratories/JanOS/common/config"
	"log"
)

func main() {
	config.Initialize()
	log.Println(`NEURON - NEUROLOGICAL SEED: ` + config.Current.Seed)

	spawnPoolSize := 16
	maxPerceptiveWidth := 36
	mm := common.NewStdMatchMaker(spawnPoolSize, maxPerceptiveWidth)

	// This will read all stdin lines and put them on a channel
	o := common.NewStdInObserver()
	go o.ForwardTo(mm.InputStream)

	common.KeepAlive()
}
