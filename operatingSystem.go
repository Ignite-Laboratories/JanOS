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
	Signals           *signalManager
	LogManager        *logManager
	RelativePath      string
	Terminate         bool
	StdBufferLength   time.Duration
	StdResolution     int
	masterCount       uint64
	worlds            []world
	terminationSignal chan os.Signal
}

// Universe The single entry point to the entire operating system.
var Universe = &operatingSystem{
	Assets:            newAssetManager(),
	Signals:           newSignalManager(),
	LogManager:        newLogManager(),
	RelativePath:      "../Assets/",
	StdBufferLength:   time.Duration(time.Second * 5),
	StdResolution:     44000,
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
func (os *operatingSystem) Start(preflight func(), realityLoop func(delta time.Duration), worlds ...world) {
	Universe.Println(os, "Hello, world")
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

	Universe.Println(os, "Starting reality loop")
	lastUpdate := time.Now()
	for {
		if Universe.Terminate {
			break
		}
		now := time.Now()
		realityLoop(now.Sub(lastUpdate))
		lastUpdate = now
		time.Sleep(1)
	}
	Universe.Println(os, "Reality loop stopped")
	Universe.Println(os, "Exiting")
}
