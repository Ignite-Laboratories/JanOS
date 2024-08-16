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

// operatingSystem Represents the core components available globally to JanOS.
type operatingSystem struct {
	Assets            *assetManager
	Dimensions        *dimensionManager
	LogManager        *logManager
	Window            *Window
	Terminate         bool
	Resolution        *resolution
	BufferLength      time.Duration
	BufferFrequency   time.Duration
	masterCount       uint64
	worlds            []world
	terminationSignal chan os.Signal
}

type resolution struct {
	Frequency   float64
	Nanoseconds int64
	Duration    time.Duration
}

// Universe The single entry point to the entire operating system.
var Universe = &operatingSystem{
	Assets:     newAssetManager(),
	Dimensions: newDimensionManager(),
	LogManager: newLogManager(),
	Resolution: &resolution{
		Frequency: 44000,
	},
	BufferLength:      time.Second * 5,
	BufferFrequency:   time.Millisecond * 10,
	worlds:            make([]world, 0),
	terminationSignal: make(chan os.Signal, 1),
}

// NextId increments the internal master count maintained since execution and then returns the value.
// This happens as an atomic operation to ensure uniqueness across threads.
func (os *operatingSystem) NextId() uint64 { return atomic.AddUint64(&os.masterCount, 1) }

// GetNamedValue returns the assigned name to this instance.
func (os *operatingSystem) GetNamedValue() string { return "JanOS" }

// Printf calls log.Print and captures the event on the os.LogManager
func (os *operatingSystem) Printf(named named, format string, v ...any) {
	os.Println(named, fmt.Sprintf(format, v...))
}

// Println calls log.Println and captures the event on the os.LogManager
func (os *operatingSystem) Println(named named, str string) {
	os.LogManager.AddEntry(named, str)
	log.Printf("[%s] %s\n", named.GetNamedValue(), str)
}

// Start initializes all the provided worlds, which initialize their systems, and then
// starts the appropriate loops to maintain the system.
func (os *operatingSystem) Start(window *Window, preflight func(), onRealityUpdate func(delta time.Duration), worlds ...world) {
	Universe.Println(os, "Hello, world")
	Universe.Printf(os, "Operating Resolution %dhz", int64(os.Resolution.Frequency))
	os.Window = window
	os.worlds = worlds
	wg := sync.WaitGroup{}

	for _, w := range os.worlds {
		wg.Add(1)
		go func() {
			if init, ok := w.(initializable); ok {
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

	realityLoop := func() {
		Universe.Println(os, "Starting reality loop")
		lastUpdate := time.Now()
		for {
			if Universe.Terminate {
				break
			}
			now := time.Now()
			onRealityUpdate(now.Sub(lastUpdate))
			os.Resolution.Nanoseconds = int64(float64(time.Second.Nanoseconds()) / os.Resolution.Frequency)
			os.Resolution.Duration = time.Duration(os.Resolution.Nanoseconds + 1)
			lastUpdate = now
			time.Sleep(1)
		}
		Universe.Println(os, "Reality loop stopped")
	}

	if os.Window != nil {
		Universe.Println(os, "Requesting a window")
		// Fork off the reality thread to give ebiten the main thread
		go realityLoop()
		os.Window.Open()
	} else {
		Universe.Println(os, "Running without a window")
		// Keep the main thread since ebiten isn't in use
		realityLoop()
	}

	Universe.Println(os, "Exiting")
}
