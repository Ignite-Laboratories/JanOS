package core

import (
	"core/sys/atlas"
	"core/sys/blue"
	"core/sys/given"
	"core/sys/given/format"
	"debug/buildinfo"
	"fmt"
	"os"
	"strings"
	"time"
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

// Alive globally keeps neural activity firing until set to false - it's true by default.
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

// WhileAlive can be used to efficiently hold a main function open.
func WhileAlive() {
	for Alive() {
		// Give the host some breathing room.
		time.Sleep(time.Millisecond)
	}
}
