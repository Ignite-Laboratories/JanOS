package Spark

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"sync/atomic"
)

/**
UTILITY FUNCTIONS
*/

var masterCount uint64

// NextId increments the internal master count maintained since execution and then returns the value
func NextId() uint64 { return atomic.AddUint64(&masterCount, 1) }

/**
ENTITY
*/

type Entity int

func NewEntity() Entity { return Entity(NextId()) }

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
	SystemIdentifier
	SystemInitializer
	SystemTicker
}

type SystemIdentifier interface {
	GetName() string
	GetEntity() Entity
}

type SystemInitializer interface {
	Initialize()
}

type SystemTicker interface {
	Tick(inbox Inbox)
}

type SystemDrawer interface {
	OnDraw(entity Entity, screen *ebiten.Image)
}

/**
WORLD
*/

type World struct {
	Messaging             Nexus
	Assets                AssetManager
	Entities              []Entity
	Systems               []System
	Cursoring             CursoringSystem
	Oscillation           OscillationSystem
	WaveformVisualization WavVizSystem
}

func (w *World) CreateEntity() Entity {
	entity := NewEntity()
	w.AddEntity(entity)
	return entity
}

func (w *World) AddEntity(e Entity) {
	w.Entities = append(w.Entities, e)
}
