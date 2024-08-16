package JanOS

import (
	"fmt"
	"time"
)

/**
ENTITY
*/

// Entity represents an ID attributable to an abstract object.
type Entity int

// NewEntity creates a new entity using the Universe's master ID counter
func NewEntity() Entity { return Entity(Universe.NextId()) }

/**
COMPONENT
*/

type Components[T any] struct {
	DB map[Entity]T
}

func (c *Components[T]) Delete(e Entity) {
	delete(c.DB, e)
}

func (c *Components[T]) Get(e Entity) (T, bool) {
	v, ok := c.DB[e]
	return v, ok
}

func (c *Components[T]) MustGet(e Entity) T {
	v, ok := c.DB[e]
	if !ok {
		panic(fmt.Sprintf("no component for entity %d", e))
	}
	return v
}

func (c *Components[T]) Set(e Entity, d T) {
	if c.DB == nil {
		c.DB = map[Entity]T{}
	}
	c.DB[e] = d
}
func (c *Components[T]) SetIfAbsent(e Entity, d T) {
	if c.DB == nil {
		c.DB = map[Entity]T{}
	}
	if _, ok := c.DB[e]; !ok {
		c.DB[e] = d
	}
}

func (c *Components[T]) Accept(entity Entity, fn func(e Entity, c T)) {
	if c, ok := c.DB[entity]; ok {
		fn(entity, c)
	}
}

// AcceptEmpty only accept if empty
func (c *Components[T]) AcceptEmpty(entity Entity, fn func(e Entity)) {
	if _, ok := c.DB[entity]; !ok {
		fn(entity)
	}
}

/**
SYSTEM
*/

type System interface {
	named
	Tick(entity Entity, delta time.Duration)
}

/**
WORLD
*/

type ecsWorld struct {
	Name     string
	Entities []Entity
	Systems  []System
}

func NewECSWorld(name string, systems ...System) world {
	return &ecsWorld{
		Name:     name,
		Entities: make([]Entity, 0),
		Systems:  systems,
	}
}

// GetNamedValue returns the assigned name to this instance.
func (w *ecsWorld) GetNamedValue() string {
	return w.Name
}

func (w *ecsWorld) Initialize() {
	for _, system := range w.Systems {
		if init, ok := system.(initializable); ok {
			Universe.Printf(w, "[%s] Initializing", system.GetNamedValue())
			init.Initialize()
		}
		Universe.Printf(w, "[%s] Initialized", system.GetNamedValue())
	}
}

func (w *ecsWorld) Start() {
	lastNow := time.Now()
	for {
		if Universe.Terminate {
			break
		}

		for _, entity := range w.Entities {
			for _, system := range w.Systems {
				system.Tick(entity, time.Since(lastNow))
			}
		}

		time.Sleep(1)
	}
}
