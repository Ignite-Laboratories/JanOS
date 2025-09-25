package deploy

import (
	"context"
	_ "embed"
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"text/template"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/sys/log"
)

//go:embed fly/fly.toml
var flyConfig string

//go:embed fly/Dockerfile
var dockerfile string

type _fly byte

var Fly _fly

// Spark will deploy the to fly.io using the given app name.  The 'target' is a path to the target main.go
// file to deploy which is relative to JanOS's root folder.  For example, to deploy the GitVanity neuron:
//
//	deploy.ToFly("appName", "navigator", "git") // Resolves to the main.go at [janOS]/navigator/git
func (_fly) Spark(flyApp string, target ...string) {
	relative := strings.Join(target, "/")
	workingDir := core.RelativePath(target...)
	depthToRoot := strings.Repeat("../", len(target))
	_root := strings.Split(workingDir, "/")
	_root = _root[:len(_root)-len(target)]
	root := strings.Join(_root, "/")
	log.Printf(ModuleName, "Sparking a deployment of '%s' to '%s'\n", relative, flyApp)

	// 0 - Create a temp folder

	_ = os.Mkdir(root+"/.tmp", 0750)
	temp, err := os.MkdirTemp(root+"/.tmp", "fly-deploy-*")
	if err != nil {
		panic(err)
	}
	defer func() {
		err = os.RemoveAll(temp)
		if err != nil {
			panic(err)
		} else {
			log.Verbosef(ModuleName, "Cleaned up temp folder '%s'\n", temp)
		}
	}()
	log.Verbosef(ModuleName, "Made temp folder '%s'\n", temp)

	// 1 - Write out the fly.toml file
	w := new(strings.Builder)

	t := template.Must(template.New("fly.toml").Parse(flyConfig))
	if err = t.Execute(w, struct{ FlyApp string }{FlyApp: flyApp}); err != nil {
		panic(err)
	}
	log.Verbosef(ModuleName, "Generated fly config:\n%v\n", w.String())

	flyPath := filepath.Join(temp, "fly.toml")
	if err = os.WriteFile(flyPath, []byte(w.String()), 0644); err != nil {
		panic(err)
	}
	log.Printf(ModuleName, "Wrote %v\n", flyPath)

	// 2 - Write out the Dockerfile
	w.Reset()

	t = template.Must(template.New("Dockerfile").Parse(dockerfile))
	if err = t.Execute(w, struct {
		Target string
		JanOS  string
		Depth  string
	}{
		Target: relative,
		JanOS:  root,
		Depth:  depthToRoot,
	}); err != nil {
		panic(err)
	}
	log.Verbosef(ModuleName, "Generated dockerfile:\n%v\n", w.String())

	dockerPath := filepath.Join(temp, "Dockerfile")
	if err = os.WriteFile(dockerPath, []byte(w.String()), 0644); err != nil {
		panic(err)
	}
	log.Printf(ModuleName, "Wrote %v\n", dockerPath)

	// 3 - Pass control to fly
	log.Printf(ModuleName, "Passing control to 'fly deploy'\n")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	deployTimeout := 33 * time.Minute
	ctx, cancel = context.WithTimeout(ctx, deployTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "fly", "deploy", depthToRoot, "--config", temp+"/"+"fly.toml")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workingDir

	if runtime.GOOS != "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	}

	if err = cmd.Start(); err != nil {
		log.FatalfCode(1, ModuleName, "failed to start fly: %v\n", err)
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
		// Preserve child’s exit code if available
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		log.FatalfCode(1, ModuleName, "deploy failed: %v\n", err)
	}

	log.Printf(ModuleName, "Successfully deployed '%v' to '%v'\n", relative, flyApp)
}
