package Logic

import (
	"JanOS"
	"fmt"
	"time"
)

/**
COMPONENT
*/

type Components[T any] struct {
	DB map[JanOS.Entity]T
}

func (c *Components[T]) Delete(e JanOS.Entity) {
	delete(c.DB, e)
}

func (c *Components[T]) Get(e JanOS.Entity) (T, bool) {
	v, ok := c.DB[e]
	return v, ok
}

func (c *Components[T]) MustGet(e JanOS.Entity) T {
	v, ok := c.DB[e]
	if !ok {
		panic(fmt.Sprintf("no component for entity %d", e))
	}
	return v
}

func (c *Components[T]) Set(e JanOS.Entity, d T) {
	if c.DB == nil {
		c.DB = map[JanOS.Entity]T{}
	}
	c.DB[e] = d
}
func (c *Components[T]) SetIfAbsent(e JanOS.Entity, d T) {
	if c.DB == nil {
		c.DB = map[JanOS.Entity]T{}
	}
	if _, ok := c.DB[e]; !ok {
		c.DB[e] = d
	}
}

func (c *Components[T]) Accept(entity JanOS.Entity, fn func(e JanOS.Entity, c T)) {
	if c, ok := c.DB[entity]; ok {
		fn(entity, c)
	}
}

// AcceptEmpty only accept if empty
func (c *Components[T]) AcceptEmpty(entity JanOS.Entity, fn func(e JanOS.Entity)) {
	if _, ok := c.DB[entity]; !ok {
		fn(entity)
	}
}

/**
SYSTEM
*/

type System interface {
	GetName() string
	Tick(entity JanOS.Entity, delta time.Duration)
}

/**
WORLD
*/

type ecsWorld struct {
	Name     string
	Entities []JanOS.Entity
	Systems  []System
}

func NewECSWorld(name string, systems ...System) JanOS.World {
	return &ecsWorld{
		Name:     name,
		Entities: make([]JanOS.Entity, 0),
		Systems:  systems,
	}
}

func (w *ecsWorld) GetName() string {
	return w.Name
}

func (w *ecsWorld) Initialize() {
	for _, system := range w.Systems {
		if init, ok := system.(JanOS.Initializable); ok {
			JanOS.Universe.Printf(w, "[%s] Initializing", system.GetName())
			init.Initialize()
		}
		JanOS.Universe.Printf(w, "[%s] Initialized", system.GetName())
	}
}

func (w *ecsWorld) Start() {
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
