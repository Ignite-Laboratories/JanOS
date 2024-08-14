package main

import (
	"JanOS"
	"JanOS/Arwen"
	"JanOS/Arwen/AI_Music"
	"JanOS/Logic"
	"JanOS/Logic/Symbol"
	"JanOS/Spark"
	"time"
)

var waveformSys = Arwen.NewWaveformSystem()

var aiMusicSys = AI_Music.NewAI_MusicSystem()
var ecsWorld = Logic.NewECSWorld("Logic", waveformSys, aiMusicSys)
var sparkWorld = Spark.NewSparkWorld("Spark")

func main() {
	JanOS.Universe.Start(preflight, tick, ecsWorld, sparkWorld)
}

var performance AI_Music.Performance
var binaryData AI_Music.BinaryData
var omega *JanOS.Dimension

func preflight() {
	performance, _ = aiMusicSys.LookupPerformance(AI_Music.FamilyBrass, AI_Music.NameTrumpetInC, AI_Music.PitchA5, AI_Music.DynamicFortissimo)
	binaryData, _ = aiMusicSys.GetBinaryData(performance.Entity)
	omega = JanOS.Universe.Dimensions.NewDimension("Omega", Symbol.Omega, 42)
}

func tick(delta time.Duration) {
}
