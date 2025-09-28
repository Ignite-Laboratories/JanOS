package neural

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/rec"
)

// SubProcess sparks off a separate process of the provided command as a neural child of the current instance.  This means
// that when the instance terminates, the child process will be cleaned up.
func (_shell) SubProcess(named string, command []string, onExit ...func(*std.Impulse)) std.Synapse {
	dir, _ := os.Getwd()
	return Shell.SubProcessAt(named, command, dir, onExit...)
}

func (_shell) SubProcessAt(named string, command []string, path string, onExit ...func(*std.Impulse)) std.Synapse {
	if len(command) == 0 {
		panic("no command provided")
	}

	cleanup := func(*std.Impulse) {}
	if len(onExit) > 0 {
		cleanup = onExit[0]
	}

	ctx, cancel := context.WithCancel(context.Background())

	return std.NewSynapse(lifecycle.Looping, named, func(imp *std.Impulse) {
		cmd := exec.CommandContext(ctx, command[0], command[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = path
		imp.Thought = std.NewThought(cmd)
		rec.Verbosef(imp.Bridge, "executing command: %v/%v\n", path, strings.Join(command, " "))
		fmt.Println(cmd.String())
		err := cmd.Start()
		if err != nil {
			panic(err)
		}

		if runtime.GOOS != "windows" {
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
		}

		go func() {
			// This goroutine is only necessary if you want explicit forwarding.
			sigCh := make(chan os.Signal, 2)
			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
			for sig := range sigCh {
				if cmd.Process != nil {
					_ = cmd.Process.Signal(sig)
				}
			}
		}()

		if err = cmd.Wait(); err != nil {
			var exitErr *exec.ExitError
			if exitErr, _ = err.(*exec.ExitError); exitErr != nil {
				rec.Printf(imp.Bridge, "[%d] sub process error %v\n", exitErr.ExitCode(), err)
			}
		}
		rec.Printf(imp.Bridge, "sub process exited\n")
		imp.Thought = nil
		cleanup(imp)
	}, func(imp *std.Impulse) bool {
		if imp.Thought == nil {
			return true
		}
		return false
	}, func(imp *std.Impulse) {
		cancel()
		ctx, cancel = context.WithCancel(context.Background())
	})
}
