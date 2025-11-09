//go:build darwin || freebsd || netbsd || openbsd

package atlas

import (
	"log"
	"os"
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

	// Open directory
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	// Create kqueue
	kq, err := syscall.Kqueue()
	if err != nil {
		f.Close()
		return nil, err
	}

	// Register directory with kqueue - watch for any changes
	ev := syscall.Kevent_t{
		Ident:  uint64(f.Fd()),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_CLEAR,
		Fflags: syscall.NOTE_WRITE | syscall.NOTE_DELETE | syscall.NOTE_EXTEND | syscall.NOTE_ATTRIB | syscall.NOTE_REVOKE,
	}

	// Initial load
	if change != nil {
		change()
	}

	go func() {
		defer syscall.Close(kq)
		defer f.Close()

		events := make([]syscall.Kevent_t, 1)
		lastReload := time.Now()
		debounceDelay := 100 * time.Millisecond

		// Set up timeout for kevent
		timeout := syscall.NsecToTimespec(50 * 1000000) // 50ms

		for {
			select {
			case <-cleanup:
				return
			default:
				n, err := syscall.Kevent(kq, []syscall.Kevent_t{ev}, events, &timeout)
				if err != nil {
					if err == syscall.EINTR {
						continue
					}
					log.Printf("Error in kevent: %v", err)
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
