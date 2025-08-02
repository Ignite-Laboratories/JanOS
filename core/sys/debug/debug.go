package debug

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetGoroutineID gets the ID of the currently executing
// goroutine by parsing it from the stack trace buffer.
//
// NOTE: There are no guarantees behind the stack trace's
// formatting!  This may or may not function going forward,
// but it's pivotal in understanding HOW execution spans
// across threads.
//
//	Works consistently on go 1.24.1
func GetGoroutineID() uint64 {
	// Get the stack trace buffer
	buf := make([]byte, 64)
	buf = buf[:runtime.Stack(buf, false)]

	// Parse the goroutine ID from the stack trace, which usually has the form "goroutine 10 [running]: ..."
	fields := bytes.Fields(buf)
	if len(fields) >= 2 {
		id, _ := strconv.ParseUint(string(fields[1]), 10, 64)
		return id
	}
	return 0
}
