//go:build linux

package atlas

import (
	"log"
	"path/filepath"
	"syscall"
	"time"
)

func watch(configPath string, onChange ...func()) (chan any, error) {
	var change func()
	if len(onChange) > 0 {
		change = onChange[0]
	}

	// Watch the directory containing the file, not the file itself
	dir := filepath.Dir(configPath)
	if dir == "" || dir == "." {
		dir = "."
	}

	// Create inotify instance
	fd, err := syscall.InotifyInit()
	if err != nil {
		return nil, err
	}

	// Make non-blocking
	syscall.SetNonblock(fd, true)

	// Watch the directory for any changes (create, modify, delete, move)
	wd, err := syscall.InotifyAddWatch(fd, dir,
		syscall.IN_CREATE|syscall.IN_MODIFY|syscall.IN_CLOSE_WRITE|
			syscall.IN_DELETE|syscall.IN_MOVED_FROM|syscall.IN_MOVED_TO)
	if err != nil {
		syscall.Close(fd)
		return nil, err
	}

	// Initial load
	if change != nil {
		change()
	}

	go func() {
		defer syscall.Close(fd)
		defer syscall.InotifyRmWatch(fd, uint32(wd))

		buf := make([]byte, syscall.SizeofInotifyEvent*10+syscall.NAME_MAX+1)
		lastReload := time.Now()
		debounceDelay := 100 * time.Millisecond

		for {
			select {
			case <-cleanup:
				return
			default:
				n, err := syscall.Read(fd, buf)
				if err != nil {
					if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK {
						// No data available, sleep briefly
						time.Sleep(50 * time.Millisecond)
						continue
					}
					log.Printf("Error reading inotify events: %v", err)
					return
				}

				if n > 0 {
					// Debounce: only reload if enough time has passed
					// The callback will check if the atlas file specifically changed
					if time.Since(lastReload) > debounceDelay {
						lastReload = time.Now()
						if change != nil {
							change()
						}
					}
				}
			}
		}
	}()

	return cleanup, nil
}
