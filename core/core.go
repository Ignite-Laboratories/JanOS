package core

import (
	"context"
	"debug/buildinfo"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"git.ignitelabs.net/janos/core/sys/atlas"
	"git.ignitelabs.net/janos/core/sys/blue"
	"git.ignitelabs.net/janos/core/sys/given"
	"git.ignitelabs.net/janos/core/sys/given/format"
)

func init() {
	exe, _ := os.Executable()
	exeInfo, _ := buildinfo.ReadFile(exe)

	if atlas.PrintPreamble {
		var version string
		for _, dep := range exeInfo.Deps {
			if dep.Path == "git.ignitelabs.net/janos/core" {
				version = dep.Version
			}
		}

		dash := "─"
		randomDash := func() string {
			return strings.Repeat(dash, int(blue.Note()))
		}

		fmt.Printf("   ╭"+randomDash()+"⇥ JanOS %v\n", version)
		fmt.Println("╭──┼" + randomDash() + "⇥ © 2025 - Humanity")
		fmt.Println("⎨  ⎬" + randomDash() + "⇥ Maintained by The Enigmaneering Guild")
		fmt.Println("╰─┬┴" + randomDash() + "⇥ ↯ [core] " + Name.StringQuoted(false))
		fmt.Println("  ╰" + randomDash() + "⬎")
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

var described = false

// Describe sets the core name's description and prints the output.  If you'd like a silent output, please
// set the Name.Description directly.
func Describe(description string) {
	Name.Description = description
	described = true
	fmt.Println("[core] " + Name.Name + " is a \"" + description + "\"")
}

// Deferrals are where you can send actions you wish to be fired just before the JanOS instance shuts down.  This is useful
// for performing global 'cleanup' operations.
func Deferrals() chan<- func(group *sync.WaitGroup) {
	return deferrals
}

var deferrals = make(chan func(*sync.WaitGroup), 1<<16)

// ShutdownLock is a part of the ShutdownCondition system.  If you would like to wait for a broadcast 'shutdown' message,
// please use ShutdownCondition.  For example:
//
//	core.ShutdownLock.Lock()
//	core.ShutdownCondition.Wait()
//	core.ShutdownLock.Unlock()
var ShutdownLock = sync.Mutex{}

// ShutdownCondition provides a way to block until the JanOS isntance broadcasts a shutdown message globally.  To do so,
// you must use the following pattern (as sync.Cond requires a sync.Mutex):
//
//	core.ShutdownLock.Lock()
//	core.ShutdownCondition.Wait()
//	core.ShutdownLock.Unlock()
var ShutdownCondition = &sync.Cond{L: &ShutdownLock}

// Shutdown waits a period of time before calling ShutdownNow.  You may optionally provide an OS exit code, otherwise
// '0' is implied.
//
// NOTE: If you don't know a proper exit code but are indicating an issue occurred, please use the catch-all exit code '1'.
func Shutdown(period time.Duration, exitCode ...int) {
	fmt.Printf("[core] %v instance shutting down in %v\n", Name.Name, period)
	time.Sleep(period)
	ShutdownNow(exitCode...)
}

// ShutdownNow immediately sets Alive to false, then pauses for a second before calling os.Exit. You may optionally
// provide an OS exit code, otherwise '0' is implied.
//
// NOTE: If you don't know a proper exit code but are indicating an issue occurred, please use the "catch-all" exit code of '1'.
func ShutdownNow(exitCode ...int) {
	fmt.Printf("\n[core] %v instance shutting down\n", Name.Name)
	alive = false

	wg := &sync.WaitGroup{}

	count := len(deferrals)
	if count > 0 {
		if count > 1 {
			fmt.Printf("[core] %v running %d deferrals\n", Name.Name, count)
		} else {
			fmt.Printf("[core] %v running %d deferral\n", Name.Name, count)
		}
		for len(deferrals) > 0 {
			deferFn := <-deferrals
			wg.Add(1)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("[core] %v deferral error: %v\n", Name.Name, r)
						wg.Done()
					}
				}()

				deferFn(wg)
			}()
		}
		wg.Wait()

		if described {
			fmt.Printf("[core] signing off — \"%v, %v\"\n", Name.Name, Name.Description)
		} else {
			fmt.Printf("[core] signing off — \"%v\"\n", Name.Name)
		}
	}

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

// KeepAlive will block the current thread until a call to Shutdown - then, this will sleep for the provided duration.
//
// NOTE: If no duration is provided, time.Duration(0) is implied.
func KeepAlive(postDelay ...time.Duration) {
	if len(postDelay) > 0 {
		deferrals <- func(wg *sync.WaitGroup) {
			fmt.Printf("[core] %v holding open for %v\n", Name.Name, postDelay[0])
			time.Sleep(postDelay[0])
			wg.Done()
		}
	}

	notify, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()
	<-notify.Done()
	ShutdownCondition.Broadcast()
	ShutdownNow()
}
