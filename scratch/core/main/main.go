// Go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.ignitelabs.net/core/sys/atlas"
)

func main() {
	// Create a context that is canceled on SIGINT (Ctrl-C) or SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Run your app (could be a server, workers, etc.) concurrently.
	done := make(chan struct{})
	go func() {
		// ... your work here ...
		<-ctx.Done() // wait until cancel requested
		// cleanup triggered by context cancellation can also happen here
		close(done)
	}()

	// Block until a signal arrives
	<-ctx.Done()
	fmt.Println("signal received, starting cleanup...")

	// Perform cleanup with a timeout so we don’t hang forever
	cleanupCtx, cancel := context.WithTimeout(context.Background(), atlas.ShutdownTimeout)
	defer cancel()
	if err := cleanup(cleanupCtx); err != nil {
		fmt.Println("cleanup error:", err)
	}

	fmt.Println("bye")
}

func cleanup(ctx context.Context) error {
	// Close DB, stop servers, flush logs, etc.
	select {
	case <-time.After(2 * time.Second):
		// pretend cleanup finished
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
