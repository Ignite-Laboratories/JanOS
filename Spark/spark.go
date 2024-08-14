package Spark

import (
	"JanOS"
	"time"
)

type SparkWorld struct {
	Name string
}

func NewSparkWorld(name string) JanOS.World {
	return &SparkWorld{
		Name: name,
	}
}

func (w *SparkWorld) GetName() string {
	return w.Name
}

func (w *SparkWorld) Initialize() {
}

func (w *SparkWorld) Start() {
	for {
		if JanOS.Universe.Terminate {
			break
		}
		time.Sleep(time.Second)
	}
}
