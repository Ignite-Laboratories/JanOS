package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/std/neural"
	"git.ignitelabs.net/janos/core/sys/rec"
)

func main() {
	cortex := std.NewCortex(std.RandomName())
	cortex.Frequency = 1 //hz
	cortex.Mute()

	port := ":4242"
	cortex.Synapses() <- neural.Net.Server("localhost"+port, port, Handler, func(imp *std.Impulse) {
		cortex.Impulse()
	})

	cortex.Spark()
	cortex.Impulse()
	core.KeepAlive(time.Second * 5)
}

func Handler(imp *std.Impulse) http.Handler {
	go func() {
		// Introduce a faux delayed shutdown

		delay := time.Second * 5
		rec.Printf(imp.Bridge, "disconnecting in %v\n", delay)
		time.Sleep(delay)
		if imp.Thought != nil && imp.Thought.Revelation != nil {
			_ = imp.Thought.Revelation.(*http.Server).Shutdown(context.Background())
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(fmt.Sprintf("\"Hello, World!\"\n\t- %v", imp.Bridge)))
	})
}
