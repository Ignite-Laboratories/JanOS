package JanOS

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// operatingSystem Represents the core components available globally
type operatingSystem struct {
	Assets            *assetManager
	Dimensions        *dimensionManager
	Window            *Window
	Terminate         bool
	Resolution        float64
	masterCount       uint64
	worlds            []World
	terminationSignal chan os.Signal
}

// Universe The single entry point to the entire operating system
var Universe = &operatingSystem{
	Assets:            newAssetManager(),
	Dimensions:        NewDimensionManager(),
	Resolution:        44000,
	worlds:            make([]World, 0),
	terminationSignal: make(chan os.Signal, 1),
}

type named interface {
	GetName() string
}

type World interface {
	named
	Start()
}

type Initializable interface {
	Initialize()
}

type Entity int

func NewEntity() Entity { return Entity(Universe.NextId()) }

// NextId increments the internal master count maintained since execution and then returns the value.
// This happens as an atomic operation to ensure uniqueness across threads.
func (os *operatingSystem) NextId() uint64 { return atomic.AddUint64(&os.masterCount, 1) }

func (os *operatingSystem) GetName() string { return "JanOS" }

func (os *operatingSystem) Printf(named named, format string, v ...any) {
	os.Println(named, fmt.Sprintf(format, v...))
}

func (os *operatingSystem) Println(named named, str string) {
	log.Printf("[%s] %s\n", named.GetName(), str)
}

func (os *operatingSystem) Start(window *Window, preflight func(), tick func(delta time.Duration), worlds ...World) {
	Universe.Println(os, "Hello, world")
	Universe.Printf(os, "Operating Resolution %dhz", int64(os.Resolution))
	os.Window = window
	os.worlds = worlds
	wg := sync.WaitGroup{}

	for _, w := range os.worlds {
		wg.Add(1)
		go func() {
			if init, ok := w.(Initializable); ok {
				Universe.Println(w, "Initializing")
				init.Initialize()
			}

			Universe.Println(w, "Starting")
			wg.Done()
			w.Start()
			// Worlds are meant to run indefinitely...
			Universe.Println(w, "Stopped")
		}()
	}
	// Wait for all the worlds to initialize
	wg.Wait()

	Universe.Println(os, "Awaiting termination signal")
	signal.Notify(os.terminationSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-os.terminationSignal
		Universe.Printf(os, "Received termination signal %s", sig)
		os.Terminate = true
	}()

	Universe.Println(os, "Running pre-flight routine")
	preflight()

	Universe.Println(os, "Taking off")

	os.Window.Open()

	Universe.Println(os, "Exiting")
}
