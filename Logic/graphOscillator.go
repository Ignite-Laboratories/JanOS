package Logic

import (
	"github.com/Ignite-Laboratories/JanOS/Logic/Math"
	"gorgonia.org/gorgonia"
	"log"
	"math"
	"time"
)

type GraphOscillator struct {
	LastUpdate     time.Time
	Result         *gorgonia.Node
	Amplitude      *gorgonia.Node
	Frequency      *gorgonia.Node
	Time           *gorgonia.Node
	LastInstant    *gorgonia.Node
	CurrentInstant *gorgonia.Node
	Graph          *gorgonia.ExprGraph
}

func NewGraphOscillator(amplitude float64, frequency float64) *GraphOscillator {
	g := gorgonia.NewGraph()

	var Pi = gorgonia.NewConstant(math.Pi, gorgonia.WithName(string(Math.Pi)))
	var Second = gorgonia.NewConstant(float64(time.Second), gorgonia.WithName("second"))
	var One = gorgonia.NewConstant(float64(1), gorgonia.WithName("1"))
	var Two = gorgonia.NewConstant(float64(2), gorgonia.WithName("2"))
	var OneEighty = gorgonia.NewConstant(float64(180), gorgonia.WithName("180"))
	var ThreeSixty = gorgonia.NewConstant(float64(360), gorgonia.WithName("360"))

	o := &GraphOscillator{
		LastUpdate:     time.Now(),
		Result:         gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName(string(Math.Y))),
		Amplitude:      gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName(string(Math.Alpha))),
		Frequency:      gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName(string(Math.Omega))),
		LastInstant:    gorgonia.NewScalar(g, gorgonia.Int64, gorgonia.WithName(string(Math.X))),
		CurrentInstant: gorgonia.NewScalar(g, gorgonia.Int64, gorgonia.WithName(string(Math.Z))),
		Time:           One,
		Graph:          g,
	}

	af1, err := gorgonia.Mul(Two, Pi)
	if err != nil {
		log.Fatal(err)
	}

	angularFrequency, err := gorgonia.Mul(af1, o.Frequency)
	if err != nil {
		log.Fatal(err)
	}

	step1, err := gorgonia.Mul(angularFrequency, o.Time)
	if err != nil {
		log.Fatal(err)
	}

	elapsedNano, err := gorgonia.Sub(o.CurrentInstant, o.LastInstant)
	if err != nil {
		log.Fatal(err)
	}

	elapsed, err := gorgonia.Div(elapsedNano, Second)
	if err != nil {
		log.Fatal(err)
	}

	ratio, err := gorgonia.Div(elapsed, Second)
	if err != nil {
		log.Fatal(err)
	}

	degrees, err := gorgonia.Mul(ratio, ThreeSixty)
	if err != nil {
		log.Fatal(err)
	}

	frequentDegrees, err := gorgonia.Mul(degrees, o.Frequency)
	if err != nil {
		log.Fatal(err)
	}

	pa1, err := gorgonia.Mul(frequentDegrees, Pi)
	if err != nil {
		log.Fatal(err)
	}

	pa2, err := gorgonia.Div(pa1, OneEighty)
	if err != nil {
		log.Fatal(err)
	}

	step2, err := gorgonia.Add(step1, pa2)
	if err != nil {
		log.Fatal(err)
	}

	step3, err := gorgonia.Sin(step2)
	if err != nil {
		log.Fatal(err)
	}

	o.Result, err = gorgonia.Mul(o.Amplitude, step3)
	if err != nil {
		log.Fatal(err)
	}

	gorgonia.Let(o.Amplitude, amplitude)
	gorgonia.Let(o.Frequency, frequency)
	gorgonia.Let(o.LastInstant, time.Now().UnixNano())

	return o
}

func (o *GraphOscillator) Tick() {
	vm := gorgonia.NewTapeMachine(o.Graph)
	defer vm.Close()

	gorgonia.Let(o.CurrentInstant, time.Now().UnixNano())

	if err := vm.RunAll(); err != nil {
		log.Fatal(err)
	}
	vm.Reset()
	log.Println(o.Result.Value())

	gorgonia.Let(o.LastInstant, o.CurrentInstant.Value())
}
