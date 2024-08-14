package Logic

import (
	"JanOS"
	"time"
)

type Entity int
type System interface {
	GetName() string
	Initialize()
	Tick(entity Entity, delta time.Duration)
}

func NewEntity() Entity { return Entity(JanOS.Universe.NextId()) }

type ECSWorld struct {
	Name     string
	Entities []Entity
	Systems  []System
}

func NewECSWorld(name string, systems ...System) JanOS.World {
	return &ECSWorld{
		Name:     name,
		Entities: make([]Entity, 0),
		Systems:  systems,
	}
}

func (w *ECSWorld) GetName() string {
	return w.Name
}

func (w *ECSWorld) Initialize() {
	for _, system := range w.Systems {
		JanOS.Universe.Printf(w, "[%s] Initializing", system.GetName())
		system.Initialize()
		JanOS.Universe.Printf(w, "[%s] Initialized", system.GetName())
	}
}

func (w *ECSWorld) Start() {
	lastNow := time.Now()
	for {
		if JanOS.Universe.Terminate {
			break
		}

		for _, entity := range w.Entities {
			for _, system := range w.Systems {
				system.Tick(entity, time.Since(lastNow))
			}
		}

		time.Sleep(time.Second)
	}
}
