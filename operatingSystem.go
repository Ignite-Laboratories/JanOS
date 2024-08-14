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
	masterCount       uint64
	Terminate         bool
	worlds            []World
	terminationSignal chan os.Signal
}

// Universe The single entry point to the entire operating system
var Universe = &operatingSystem{
	Assets:            newAssetManager(),
	worlds:            make([]World, 0),
	terminationSignal: make(chan os.Signal, 1),
}

type Named interface {
	GetName() string
}

type World interface {
	Named
	Initialize()
	Start()
}

// NextId increments the internal master count maintained since execution and then returns the value.
// This happens as an atomic operation to ensure uniqueness across threads.
func (os *operatingSystem) NextId() uint64 { return atomic.AddUint64(&os.masterCount, 1) }

func (os *operatingSystem) GetName() string { return "JanOS" }

func (os *operatingSystem) Printf(named Named, format string, v ...any) {
	os.Println(named, fmt.Sprintf(format, v...))
}

func (os *operatingSystem) Println(named Named, str string) {
	log.Printf("[%s] %s", named.GetName(), str)
}

func (os *operatingSystem) Start(tick func(delta time.Duration), worlds ...World) {
	Universe.Println(os, "Hello, world")
	os.worlds = worlds
	wg := sync.WaitGroup{}

	for _, w := range os.worlds {
		wg.Add(1)
		go func() {
			Universe.Println(w, "Initializing")
			w.Initialize()
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

	Universe.Println(os, "Starting Reality Loop")
	lastUpdate := time.Now()
	for {
		if os.Terminate {
			break
		}
		tick(time.Since(lastUpdate))
		time.Sleep(time.Nanosecond)
	}

	Universe.Println(os, "Exiting")
}
