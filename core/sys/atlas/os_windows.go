//go:build windows

package atlas

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows"
)

func watch(configPath string, onChange ...func()) (chan any, error) {
	var change func()
	if len(onChange) > 0 {
		change = onChange[0]
	}

	// Resolve directory to watch (treat file paths implicitly).
	dir := configPath
	if dir == "" {
		dir = "."
	}
	if fi, err := os.Stat(dir); err == nil {
		if !fi.IsDir() {
			dir = filepath.Dir(dir)
		}
	} else {
		// If the path doesn't exist yet, watch its parent directory (or ".")
		if d := filepath.Dir(dir); d != "" && d != "." && d != string(filepath.Separator) {
			dir = d
		} else {
			dir = "."
		}
	}

	const filter = windows.FILE_NOTIFY_CHANGE_FILE_NAME |
		windows.FILE_NOTIFY_CHANGE_DIR_NAME |
		windows.FILE_NOTIFY_CHANGE_ATTRIBUTES |
		windows.FILE_NOTIFY_CHANGE_SIZE |
		windows.FILE_NOTIFY_CHANGE_LAST_WRITE |
		windows.FILE_NOTIFY_CHANGE_CREATION |
		windows.FILE_NOTIFY_CHANGE_SECURITY

	h, err := windows.FindFirstChangeNotification(dir, false, filter)
	if err != nil {
		return nil, err
	}

	// Manual-reset event used to signal cancellation.
	cancelEvent, err := windows.CreateEventEx(
		nil,                                            // security attributes
		nil,                                            // name
		windows.CREATE_EVENT_MANUAL_RESET,              // manual reset, initially not signaled
		windows.EVENT_MODIFY_STATE|windows.SYNCHRONIZE, // desired access
	)
	if err != nil {
		_ = windows.FindCloseChangeNotification(h)
		return nil, err
	}

	cleanup := make(chan any)

	// Initial load before entering the loop.
	if change != nil {
		change()
	}

	// Bridge the cleanup channel to the cancellation event.
	go func() {
		<-cleanup // any value (commonly nil) triggers shutdown
		_ = windows.SetEvent(cancelEvent)
	}()

	go func() {
		defer func() {
			_ = windows.FindCloseChangeNotification(h)
			_ = windows.CloseHandle(cancelEvent)
		}()

		handles := []windows.Handle{h, cancelEvent}
		debounce := 100 * time.Millisecond
		last := time.Now().Add(-debounce)

		for {
			status, err := windows.WaitForMultipleObjects(handles, false, windows.INFINITE)
			if err != nil {
				log.Printf("watch: WaitForMultipleObjects error: %v", err)
				return
			}

			switch int(status - windows.WAIT_OBJECT_0) {
			case 0: // directory change
				if time.Since(last) >= debounce {
					last = time.Now()
					if change != nil {
						change()
					}
				}
				if err := windows.FindNextChangeNotification(h); err != nil {
					log.Printf("watch: FindNextChangeNotification error: %v", err)
					return
				}
			case 1: // cancel
				return
			default:
				// Unexpected; exit gracefully.
				return
			}
		}
	}()

	return cleanup, nil
}
