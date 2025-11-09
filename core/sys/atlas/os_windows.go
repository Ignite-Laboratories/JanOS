//go:build windows

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

	// Watch the directory containing the file
	dir := filepath.Dir(configPath)
	if dir == "" || dir == "." {
		dir = "."
	}

	// Open directory
	dirPath, err := syscall.UTF16PtrFromString(dir)
	if err != nil {
		return nil, err
	}

	handle, err := syscall.CreateFile(
		dirPath,
		syscall.FILE_LIST_DIRECTORY,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_FLAG_BACKUP_SEMANTICS|syscall.FILE_FLAG_OVERLAPPED,
		0,
	)
	if err != nil {
		return nil, err
	}

	// Initial load
	if change != nil {
		change()
	}

	go func() {
		defer syscall.CloseHandle(handle)

		buf := make([]byte, 4096)
		lastReload := time.Now()
		debounceDelay := 100 * time.Millisecond

		// Create event for overlapped I/O
		event, _ := syscall.CreateEvent(nil, 0, 0, nil)
		defer syscall.CloseHandle(event)

		overlapped := &syscall.Overlapped{HEvent: event}

		for {
			select {
			case <-cleanup:
				return
			default:
				var bytesReturned uint32
				err := syscall.ReadDirectoryChanges(
					handle,
					&buf[0],
					uint32(len(buf)),
					false,
					syscall.FILE_NOTIFY_CHANGE_FILE_NAME|syscall.FILE_NOTIFY_CHANGE_LAST_WRITE|syscall.FILE_NOTIFY_CHANGE_SIZE,
					&bytesReturned,
					overlapped,
					0,
				)

				if err != nil && err != syscall.ERROR_IO_PENDING {
					log.Printf("Error watching directory: %v", err)
					return
				}

				// Wait with timeout
				result, _ := syscall.WaitForSingleObject(event, 50) // 50ms timeout

				if result == syscall.WAIT_OBJECT_0 {
					// Get actual bytes read
					syscall.GetOverlappedResult(handle, overlapped, &bytesReturned, false)

					if bytesReturned > 0 {
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
		}
	}()

	return cleanup, nil
}
