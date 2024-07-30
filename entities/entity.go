package entities

var lastEntity int

type Entity int

func NewEntity() Entity {
	lastEntity++
	return Entity(lastEntity)
}
