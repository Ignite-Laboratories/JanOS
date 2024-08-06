package Systems

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
)

type InputSystem struct {
	Logic.Entity

	components InputSystemComponents
}

type InputSystemComponents struct {
}

func NewInputSystem() InputSystem {
	return InputSystem{
		components: InputSystemComponents{},
	}
}

func (sys InputSystem) GetName() string         { return "Input" }
func (sys InputSystem) GetEntity() Logic.Entity { return sys.Entity }

func (sys InputSystem) Initialize(world *Logic.World) {
}

var i = 0

type RawrrrMessage struct {
	Logic.Entity
	Message string
}

type OtherMessage struct {
	Logic.Entity
	Message string
}

func (sys InputSystem) Tick(world *Logic.World, inbox Logic.Inbox) {
	i++

	if i > 100 {
		world.Publish("RAWRRR", &RawrrrMessage{Message: "Rawrrr 1"})
		world.Publish("RAWRRR", &OtherMessage{Message: "Rawrrr 2"})
		world.Publish("RAWRRR", &RawrrrMessage{Message: "Rawrr 3"})
		world.Publish("Other", &RawrrrMessage{Message: "Other 1"})
		i = 0
	}
}
