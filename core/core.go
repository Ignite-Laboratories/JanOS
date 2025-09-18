package core

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"git.ignitelabs.net/core/sys/atlas"
	"git.ignitelabs.net/core/sys/blue"
	"git.ignitelabs.net/core/sys/given"
	"git.ignitelabs.net/core/sys/given/format"
)

func init() {
	exe, _ := os.Executable()
	exeInfo, _ := buildinfo.ReadFile(exe)

	if atlas.PrintPreamble {
		var version string
		for _, dep := range exeInfo.Deps {
			if strings.Contains(dep.Path, "github.com/ignite-laboratories/core") {
				version = dep.Version
			}
		}

		randDash := func() string {
			var builder strings.Builder
			for i := byte(0); i < blue.Note(); i++ {
				builder.WriteString("-")
			}
			return builder.String()
		}

		fmt.Printf("⎧"+randDash()+"-⇥ JanOS %v\n", version)
		fmt.Println("|" + randDash() + "-⇥ © 2025, Ignite Laboratories")
		fmt.Println("⎨" + randDash() + "-⇥ Alex Petz")
		fmt.Println("|" + randDash() + "⇥")
		fmt.Println("⎩" + randDash() + "⇥ ↯ " + Name.StringQuoted(false))
		fmt.Println()
	}
}

var ModuleName = "core"

// Alive indicates if the system is still alive - all JanOS instances are alive at creation.
func Alive() bool {
	return alive
}

var alive = true

// Inception provides the moment this operating system was initialized.
var Inception = time.Now()

// Name provides the randomly selected name of this instance.
var Name, _ = given.Random[format.Default]()

// Shutdown waits a period of time before calling ShutdownNow.  You may optionally provide an OS exit code, otherwise
// '0' is implied.
//
// NOTE: If you don't know a proper exit code but are indicating an issue occurred, please use the catch-all exit code '1'.
func Shutdown(period time.Duration, exitCode ...int) {
	fmt.Sprintf("[core] shutting down in %v\n", period)
	time.Sleep(period)
	ShutdownNow(exitCode...)
}

// ShutdownNow immediately sets Alive to false, then pauses for a second before calling os.Exit. You may optionally
// provide an OS exit code, otherwise '0' is implied.
//
// NOTE: If you don't know a proper exit code but are indicating an issue occurred, please use the catch-all exit code '1'.
func ShutdownNow(exitCode ...int) {
	fmt.Sprintf("[core] shutting down\n")
	alive = false

	// Give the threads a brief moment to clean themselves up.
	time.Sleep(time.Second)
	if len(exitCode) > 0 {
		os.Exit(exitCode[0])
	} else {
		os.Exit(0)
	}
}

// Ignite begins execution of NeuralActivity.  This is a blocking call, meaning your initial neural sparks should
// be provided to the channel before invoking this - but further neural activity can be sparked from any thread.
func Ignite() {
	defer func() {
		for fn := range Defer {
			fn()
		}
	}()

	for Alive() {

	}
}

type Context struct {
	LastActivation time.Time
	Now            time.Time
	Beat           int
}

// NeuralActivity sparks neural execution within a stable isolated container. This will wrap and invoke each provided
// neural Spark as a goroutine within a panic-safe environment, printing any neural panics that bubble up.
//
// NOTE: neural activity does not begin until a call to Ignite.
var NeuralActivity chan Spark = make(chan Spark, 1)

// Defer is where you can send actions you wish to be fired just before the cortex shuts down.  This is useful
// for performing 'cleanup' operations when another neuron has requested a shutdown event.
var Defer chan func() = make(chan func(), 1)

// A Neuron is any type that can fire an action-potential (see.ActionPotentials).  The Cleanup function
// will be called whenever the neuron's Lifecycle has reached completion or the cortex shuts down.
type Neuron interface {
	Action(Context)
	Potential(Context) bool
	Cleanup()
}

// A Spark initiates a Lifecycle-bound request to activate a Neuron.
type Spark struct {
	Lifecycle
	Neurons []*Neuron
}

func NewSpark(lifecycle Lifecycle, neurons ...*Neuron) Spark {
	return Spark{
		Lifecycle: lifecycle,
		Neurons:   neurons,
	}
}

// A Lifecycle defines how the neuron should be re-activated.  When a neuron is 'activated' it's provided a
// -new- goroutine, so please be wary of that.  There are four ways a neuron can be activated:
//
// 0 - Looping - this will cyclically re-activate the neuron only after it finishes its current execution
//
// 1 - Stimulative - this will activate the neuron whenever its potential returns high, regardless of current execution
//
// 2 - Triggered - this will perform a single impulse once the potential goes high.
//
// 3 - Impulse - this will ATTEMPT to activate the neuron -exactly- once.
type Lifecycle byte

const (
	Looping Lifecycle = iota
	Stimulative
	Triggered
	Impulse
)

// RelativePath returns a relative path from the root of the encapsulating JanOS directory.  You may optionally
// provide components to append to the path.  For example:
//
//	RelativePath("navigator", "git")
//
// Would return a path to the "Git Vanity URL" cortex like this:
//
//	/users/ignite/source/janOS/navigator/git
func RelativePath(components ...string) string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(file)) + "/" + strings.Join(components, "/")
}
