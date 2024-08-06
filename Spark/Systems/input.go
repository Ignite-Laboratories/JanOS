package Systems

import (
	"github.com/Ignite-Laboratories/JanOS/Spark"
)

type InputSystem struct {
	Spark.Entity

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
func (sys InputSystem) GetEntity() Spark.Entity { return sys.Entity }

func (sys InputSystem) Initialize() {
}

func (sys InputSystem) Tick(inbox Spark.Inbox) {
}
