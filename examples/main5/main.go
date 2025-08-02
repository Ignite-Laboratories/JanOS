package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func main() {
	fmt.Println("X11 Mouse Position Example")

	// Connect to the X server
	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("Failed to connect to X server: %v", err)
	}
	defer conn.Close()

	// Get the root window of the default screen
	setup := xproto.Setup(conn)
	root := setup.DefaultScreen(conn).Root

	// Infinite loop to track the mouse position
	for {
		// Query the pointer to get global mouse position
		reply, err := xproto.QueryPointer(conn, root).Reply()
		if err != nil {
			log.Printf("Failed to query pointer: %v", err)
			continue
		}

		// Print the global mouse position
		fmt.Printf("Global Mouse Position: X=%d, Y=%d\n", reply.RootX, reply.RootY)

		// Sleep briefly to avoid flooding the terminal
		time.Sleep(500 * time.Millisecond)
	}
}
