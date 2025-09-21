// Go
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func StartHTTP(addr string, h http.Handler) (actualAddr string, stop func(context.Context) error, err error) {
	ln, err := net.Listen("tcp", addr) // bind first so startup errors are handled synchronously
	if err != nil {
		return "", nil, err
	}

	srv := &http.Server{
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		// Optional: BaseContext to attach app context to connections
		// BaseContext: func(net.Listener) context.Context { return ctx },
	}

	errCh := make(chan error, 1)
	go func() {
		// Serve blocks until listener closed or server shut down
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
		close(errCh)
	}()

	// stop gracefully
	stop = func(ctx context.Context) error {
		// Shutdown stops accepting new conns and gracefully closes existing ones.
		// If you want to hard-stop, use srv.Close().
		return srv.Shutdown(ctx)
	}

	return ln.Addr().String(), stop, nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hi")
	})

	addr, stop, err := StartHTTP(":0", mux) // :0 picks a free port
	if err != nil {
		panic(err)
	}
	fmt.Println("listening on", addr)

	// ... your app runs without being blocked ...

	// When shutting down:
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = stop(ctx)
}
