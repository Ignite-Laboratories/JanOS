package components

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/entities"
)

type Components[T any] struct {
	db map[entities.Entity]T
}

func (c *Components[T]) Del(e entities.Entity) {
	delete(c.db, e)
}

func (c *Components[T]) Get(e entities.Entity) (T, bool) {
	v, ok := c.db[e]
	return v, ok
}

func (c *Components[T]) MustGet(e entities.Entity) T {
	v, ok := c.db[e]
	if !ok {
		panic(fmt.Sprintf("no component for entity %d", e))
	}
	return v
}

func (c *Components[T]) Set(e entities.Entity, d T) {
	if c.db == nil {
		c.db = map[entities.Entity]T{}
	}
	c.db[e] = d
}
func (c *Components[T]) SetIfAbsent(e entities.Entity, d T) {
	if c.db == nil {
		c.db = map[entities.Entity]T{}
	}
	if _, ok := c.db[e]; !ok {
		c.db[e] = d
	}
}

func (c *Components[T]) Accept(entity entities.Entity, fn func(e entities.Entity, c T)) {
	if c, ok := c.db[entity]; ok {
		fn(entity, c)
	}
}

// AcceptEmpty only accept if empty
func (c *Components[T]) AcceptEmpty(entity entities.Entity, fn func(e entities.Entity)) {
	if _, ok := c.db[entity]; !ok {
		fn(entity)
	}
}
