package neural

import (
	"fmt"
	"net/http"

	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
)

func (_net) HelloWorld(lifecycle lifecycle.Lifecycle, named string, address string, onDisconnect ...func(*std.Impulse)) std.Synapse {
	return Net.Server(lifecycle, named, address, func(imp *std.Impulse) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(fmt.Sprintf("\"Hello, World!\"\n\t— %v", imp.Bridge)))
		})
	}, onDisconnect...)
}
