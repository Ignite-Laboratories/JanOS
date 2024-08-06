package Systems

import (
	"github.com/Ignite-Laboratories/JanOS/Spark"
)

type OscillationSystem struct {
	Spark.Entity

	components OscillationSystemComponents
}

type OscillationSystemComponents struct {
	BinaryData *Spark.BinaryDataSet
}

func NewAudioSystem() OscillationSystem {
	return OscillationSystem{
		components: OscillationSystemComponents{
			BinaryData: &Spark.BinaryDataSet{},
		},
	}
}

func (sys OscillationSystem) GetName() string         { return "Oscillation" }
func (sys OscillationSystem) GetEntity() Spark.Entity { return sys.Entity }

func (sys OscillationSystem) Initialize() {
	Spark.Universe.Subscribe(sys.GetEntity(), "RAWRRR")
	Spark.Universe.Subscribe(sys.GetEntity(), "Other")
}

func (sys OscillationSystem) Tick(inbox Spark.Inbox) {
}
