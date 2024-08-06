package Systems

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Common"
)

type OscillationSystem struct {
	Logic.Entity

	components OscillationSystemComponents
}

type OscillationSystemComponents struct {
	BinaryData *Common.BinaryDataSet
}

func NewAudioSystem() OscillationSystem {
	return OscillationSystem{
		components: OscillationSystemComponents{
			BinaryData: &Common.BinaryDataSet{},
		},
	}
}

func (sys OscillationSystem) GetName() string         { return "Oscillation" }
func (sys OscillationSystem) GetEntity() Logic.Entity { return sys.Entity }

func (sys OscillationSystem) Initialize(world *Logic.World) {
	world.Subscribe(sys.GetEntity(), "RAWRRR")
	world.Subscribe(sys.GetEntity(), "Other")
}

func (sys OscillationSystem) Tick(world *Logic.World, inbox Logic.Inbox) {
}
