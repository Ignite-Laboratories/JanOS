package AI_Music

import (
	"JanOS"
	"JanOS/Logic"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/go-audio/wav"
	"log"
	"os"
	"strconv"
	"time"
)

type ai_musicSystem struct {
	Entries []Performance
}

type Performance struct {
	Path             string
	PathID           int
	Family           Family
	NameAbbr         string
	Name             InstrumentName
	TechniqueAbbr    string
	Technique        Technique
	Pitch            Pitch
	PitchID          int
	Dynamic          Dynamic
	DynamicID        int
	InstanceID       int
	StringID         int
	DigitallyReTuned bool
	Asset            JanOS.Asset
}

func (sys *ai_musicSystem) LookupPerformance(family Family, name InstrumentName, pitch Pitch, dynamic Dynamic) Performance {
	var performance Performance
	for _, v := range sys.Entries {
		if v.Family == family && v.Name == name && v.Pitch == pitch && v.Dynamic == dynamic {
			performance = v
			assetName := fmt.Sprintf("AI_Music.Audio.%s.%s.%s.%s", family, name, pitch, dynamic)
			JanOS.Universe.Assets.LoadAsset(assetName, fmt.Sprintf("./Assets/Training/AI_Music/%s", performance.Path))
			asset := JanOS.Universe.Assets.GetAsset(assetName)
			performance.Asset = asset
			return performance
		}
	}
	panic("No performance found!")
}

func NewAI_MusicSystem() *ai_musicSystem {
	return &ai_musicSystem{}
}

func (sys *ai_musicSystem) Initialize() {
	metadataName := "AI_Music.Audio.Metadata"
	JanOS.Universe.Assets.LoadAsset(metadataName, "./Assets/Training/AI_Music/audio.csv")
	metadataAsset := JanOS.Universe.Assets.GetAsset(metadataName)

	csvReader := csv.NewReader(bytes.NewReader(metadataAsset.Data.([]byte)))
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Skip the header
	records = records[1:]

	performances := make([]Performance, len(records))
	for i, r := range records {
		pathID, _ := strconv.Atoi(r[1])
		pitchID, _ := strconv.Atoi(r[8])
		dynamicID, _ := strconv.Atoi(r[10])
		instanceID, _ := strconv.Atoi(r[11])
		stringID, _ := strconv.Atoi(r[12])
		digitallyReTuned, _ := strconv.ParseBool(r[13])

		performances[i] = Performance{
			Path:             r[0],
			PathID:           pathID,
			Family:           Family(r[2]),
			NameAbbr:         r[3],
			Name:             InstrumentName(r[4]),
			TechniqueAbbr:    r[5],
			Technique:        Technique(r[6]),
			Pitch:            Pitch(r[7]),
			PitchID:          pitchID,
			Dynamic:          Dynamic(r[9]),
			DynamicID:        dynamicID,
			InstanceID:       instanceID,
			StringID:         stringID,
			DigitallyReTuned: digitallyReTuned,
		}
	}
	sys.Entries = performances
}

func (sys *ai_musicSystem) Tick(entity Logic.Entity, delta time.Duration) {

}

func (sys *ai_musicSystem) GetName() string { return "AI_Music System" }

func LoadWaveform(path string) []int {
	f, err := os.Open("./data/audio/" + path)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	d := wav.NewDecoder(f)
	pcmData, err := d.FullPCMBuffer()
	if err != nil {
		log.Panic(err)
	}

	return pcmData.Data
}

type Family string
type Technique string
type InstrumentName string
type Pitch string
type Dynamic string

const (
	NameBass          InstrumentName = "Bass"
	NameTuba          InstrumentName = "Tuba"
	NameFrench        InstrumentName = "French"
	NameHorn          InstrumentName = "Horn"
	NameTrombone      InstrumentName = "Trombone"
	NameTrumpetInC    InstrumentName = "Trumpet in C"
	NameAccordion     InstrumentName = "Accordion"
	NameCello         InstrumentName = "Cello"
	NameContrabass    InstrumentName = "Contrabass"
	NameViola         InstrumentName = "Viola"
	NameViolin        InstrumentName = "Violin"
	NameAltoSaxophone InstrumentName = "Alto Saxophone"
	NameBassoon       InstrumentName = "Bassoon"
	NameClarinetInBb  InstrumentName = "Clarinet in Bb"
	NameFlute         InstrumentName = "Flute"
	NameOboe          InstrumentName = "Oboe"
)

const (
	PitchFs1 Pitch = "F#1"
	PitchG1  Pitch = "G1"
	PitchGs1 Pitch = "G#1"
	PitchA1  Pitch = "A1"
	PitchAs1 Pitch = "A#1"
	PitchB1  Pitch = "B1"
	PitchC2  Pitch = "C2"
	PitchCs2 Pitch = "C#2"
	PitchD2  Pitch = "D2"
	PitchDs2 Pitch = "D#2"
	PitchE2  Pitch = "E2"
	PitchF2  Pitch = "F2"
	PitchFs2 Pitch = "F#2"
	PitchG2  Pitch = "G2"
	PitchGs2 Pitch = "G#2"
	PitchA2  Pitch = "A2"
	PitchAs2 Pitch = "A#2"
	PitchB2  Pitch = "B2"
	PitchC3  Pitch = "C3"
	PitchCs3 Pitch = "C#3"
	PitchD3  Pitch = "D3"
	PitchDs3 Pitch = "D#3"
	PitchE3  Pitch = "E3"
	PitchF3  Pitch = "F3"
	PitchFs3 Pitch = "F#3"
	PitchG3  Pitch = "G3"
	PitchGs3 Pitch = "G#3"
	PitchA3  Pitch = "A3"
	PitchAs3 Pitch = "A#3"
	PitchB3  Pitch = "B3"
	PitchC4  Pitch = "C4"
	PitchCs4 Pitch = "C#4"
	PitchD4  Pitch = "D4"
	PitchDs4 Pitch = "D#4"
	PitchE4  Pitch = "E4"
	PitchF4  Pitch = "F4"
	PitchFs4 Pitch = "F#4"
	PitchG4  Pitch = "G4"
	PitchGs4 Pitch = "G#4"
	PitchA4  Pitch = "A4"
	PitchAs4 Pitch = "A#4"
	PitchB4  Pitch = "B4"
	PitchC5  Pitch = "C5"
	PitchCs5 Pitch = "C#5"
	PitchD5  Pitch = "D5"
	PitchDs5 Pitch = "D#5"
	PitchE5  Pitch = "E5"
	PitchF5  Pitch = "F5"
	PitchFs5 Pitch = "F#5"
	PitchG5  Pitch = "G5"
	PitchGs5 Pitch = "G#5"
	PitchA5  Pitch = "A5"
	PitchAs5 Pitch = "A#5"
	PitchB5  Pitch = "B5"
	PitchC6  Pitch = "C6"
	PitchCs6 Pitch = "C#6"
	PitchD6  Pitch = "D6"
	PitchE1  Pitch = "E1"
	PitchF1  Pitch = "F1"
	PitchDs6 Pitch = "D#6"
	PitchE6  Pitch = "E6"
	PitchF6  Pitch = "F6"
	PitchFs6 Pitch = "F#6"
	PitchG6  Pitch = "G6"
	PitchGs6 Pitch = "G#6"
	PitchA6  Pitch = "A6"
	PitchAs6 Pitch = "A#6"
	PitchB6  Pitch = "B6"
	PitchC7  Pitch = "C7"
	PitchCs7 Pitch = "C#7"
	PitchD7  Pitch = "D7"
	PitchDs7 Pitch = "D#7"
	PitchE7  Pitch = "E7"
	PitchF7  Pitch = "F7"
	PitchFs7 Pitch = "F#7"
	PitchG7  Pitch = "G7"
	PitchGs7 Pitch = "G#7"
	PitchA7  Pitch = "A7"
	PitchAs7 Pitch = "A#7"
	PitchB7  Pitch = "B7"
	PitchC8  Pitch = "C8"
	PitchCs8 Pitch = "C#8"
)

const (
	DynamicPianissimo Dynamic = "pp"
	DynamicPiano      Dynamic = "p"
	DynamicMezzoForte Dynamic = "mf"
	DynamicFortissimo Dynamic = "ff"
)

const (
	FamilyBrass     Family = "Brass"
	FamilyKeyboards Family = "Keyboards"
	FamilyStrings   Family = "Strings"
	FamilyWinds     Family = "Winds"
)

const (
	TechniqueOrdinario Technique = "ordinario"
)
