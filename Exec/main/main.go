package main

import (
	"JanOS"
	"JanOS/Arwen"
	"JanOS/Arwen/AI_Music"
	"JanOS/Logic"
	"JanOS/Spark"
	"time"
)

var waveformSys = Arwen.NewWaveformSystem()

var aiMusicSys = AI_Music.NewAI_MusicSystem()
var ecsWorld = Logic.NewECSWorld("Logic", waveformSys, aiMusicSys)
var sparkWorld = Spark.NewSparkWorld("Spark")

func main() {
	JanOS.Universe.Start(tick, ecsWorld, sparkWorld)
}

var test bool

func tick(delta time.Duration) {
	if !test {
		p := aiMusicSys.LookupPerformance(AI_Music.FamilyBrass, AI_Music.NameTrumpetInC, AI_Music.PitchA5, AI_Music.DynamicFortissimo)
		JanOS.Universe.Println(aiMusicSys, string(p.Name))
		test = true
	}
}
