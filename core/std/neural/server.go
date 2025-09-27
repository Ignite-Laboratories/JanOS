package neural

import (
	"context"
	"errors"
	"net/http"

	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/rec"
)

func (_net) Server(named string, address string, handlerFn func(imp *std.Impulse) http.Handler, onDisconnect ...func(*std.Impulse)) std.Synapse {
	c := func(*std.Impulse) {}
	if len(onDisconnect) > 0 {
		c = onDisconnect[0]
	}
	if handlerFn == nil {
		panic(errors.New("handler function is nil"))
	}

	return std.NewSynapse(lifecycle.Looping, named, func(imp *std.Impulse) {
		server := &http.Server{
			Addr:    address,
			Handler: handlerFn(imp),
		}

		imp.Thought = std.NewThought(server)

		go func() {
			rec.Verbosef(imp.Bridge, "neural server listening on %s\n", address)

			if err := server.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					rec.Printf(imp.Bridge, "%s\n", err)
				} else {
					rec.Printf(imp.Bridge, "neural server error: %s\n", err)
				}
				if onDisconnect != nil {
					c(imp)
				}
			}

			imp.Thought = nil
		}()
	}, func(imp *std.Impulse) bool {
		if imp.Thought == nil {
			return true
		}
		return false
	}, func(imp *std.Impulse) {
		if imp.Thought != nil {
			imp.Thought.Gate.Lock()
			defer imp.Thought.Gate.Unlock()

			_ = imp.Thought.Revelation.(*http.Server).Shutdown(context.Background())
		}
	})
}
