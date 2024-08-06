package Systems

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Common"
	"log"
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

func ProcessMessages(inbox Logic.Inbox, rawrrrMessages func(msg *RawrrrMessage), otherMessages func(msg *OtherMessage)) {
	for subject, messages := range inbox.Subjects {
		log.Printf("[%d] %s messages received", len(messages), subject)
		for _, message := range messages {
			switch m := message.(type) {
			case *RawrrrMessage:
				rawrrrMessages(m)
			}
		}
	}
}

func (sys OscillationSystem) Tick(world *Logic.World, inbox Logic.Inbox) {
	for subject, messages := range inbox.Subjects {
		log.Printf("%d %s messages received", len(messages), subject)
		for _, message := range messages {
			switch m := message.(type) {
			case *RawrrrMessage:
				log.Print("Rawrrr: " + m.Message)
			case *OtherMessage:
				log.Print("Other: " + m.Message)

				for _, system := range world.Systems {
					switch s := system.(type) {
					case AssetSystem:
						s.LoadFile("segoe-print", "fonts\\segoepr.ttf", world)
					}
				}
			}
		}
	}
}
