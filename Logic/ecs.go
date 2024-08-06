package Logic

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

/**
ENTITY
*/

type Entity int

func NewEntity() Entity { return Entity(NextId()) }

/**
COMPONENT
*/

type Components[T any] struct {
	db map[Entity]T
}

func (c *Components[T]) Delete(e Entity) {
	delete(c.db, e)
}

func (c *Components[T]) Get(e Entity) (T, bool) {
	v, ok := c.db[e]
	return v, ok
}

func (c *Components[T]) MustGet(e Entity) T {
	v, ok := c.db[e]
	if !ok {
		panic(fmt.Sprintf("no component for entity %d", e))
	}
	return v
}

func (c *Components[T]) Set(e Entity, d T) {
	if c.db == nil {
		c.db = map[Entity]T{}
	}
	c.db[e] = d
}
func (c *Components[T]) SetIfAbsent(e Entity, d T) {
	if c.db == nil {
		c.db = map[Entity]T{}
	}
	if _, ok := c.db[e]; !ok {
		c.db[e] = d
	}
}

func (c *Components[T]) Accept(entity Entity, fn func(e Entity, c T)) {
	if c, ok := c.db[entity]; ok {
		fn(entity, c)
	}
}

// AcceptEmpty only accept if empty
func (c *Components[T]) AcceptEmpty(entity Entity, fn func(e Entity)) {
	if _, ok := c.db[entity]; !ok {
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
	Initialize(world *World)
}

type SystemTicker interface {
	Tick(world *World, inbox Inbox)
}

type SystemDrawer interface {
	Draw(img *ebiten.Image)
}

/**
WORLD
*/

type World struct {
	Messaging Nexus
	Assets    AssetManager
	Entities  []Entity
	Systems   []System
}

func (w *World) AddEntity(e Entity) {
	w.Entities = append(w.Entities, e)
}
