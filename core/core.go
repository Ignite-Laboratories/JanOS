package core

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
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

		randomDash := func() string {
			var builder strings.Builder
			for i := byte(0); i < blue.Note(); i++ {
				builder.WriteString("-")
			}
			return builder.String()
		}

		fmt.Printf("⎧"+randomDash()+"-⇥ JanOS %v\n", version)
		fmt.Println("|" + randomDash() + "-⇥ © 2025, Ignite Laboratories")
		fmt.Println("⎨" + randomDash() + "-⇥ Alex Petz")
		fmt.Println("|" + randomDash() + "⇥")
		fmt.Println("⎩" + randomDash() + "⇥ ↯ " + Name.StringQuoted(false))
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
var Name = given.Random[format.Default]()

// Deferrals are where you can send actions you wish to be fired just before the JanOS instance shuts down.  This is useful
// for performing global 'cleanup' operations.
func Deferrals() chan<- func() {
	return deferrals
}

var deferrals = make(chan func(), 2^16)

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
// NOTE: If you don't know a proper exit code but are indicating an issue occurred, please use the "catch-all" exit code of '1'.
func ShutdownNow(exitCode ...int) {
	fmt.Sprintf("[core] shutting down\n")
	alive = false

	wg := sync.WaitGroup{}
	for deferFn := range deferrals {
		wg.Add(1)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("[%s] deferral error: %v\n", ModuleName, r)
					wg.Done()
				}
			}()

			deferFn()
			wg.Done()
		}()
	}
	wg.Wait()

	// Give the threads a brief moment to clean themselves up.
	time.Sleep(time.Second)
	if len(exitCode) > 0 {
		os.Exit(exitCode[0])
	} else {
		os.Exit(0)
	}
}

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
