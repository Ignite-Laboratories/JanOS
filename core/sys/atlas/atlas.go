package atlas

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"sync"
)

var keys = make(map[string]any)
var gate sync.Mutex
var cleanup = make(chan any)

func init() {
	_, _ = watch(".", refresh)
}

func refresh() {
	gate.Lock()
	defer gate.Unlock()

	path := "atlas"

	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	if len(data) == 0 {
		return
	}

	var cfg config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return
	}
	cfg.apply()

	if err = json.Unmarshal(data, &keys); err != nil {
		return
	}
}

// Cleanup is called by core on shutdown to ensure the file watchers are closed.
func Cleanup() {
	if cleanup != nil {
		cleanup <- nil
	}
}

// Parse will read the atlas file and attempt to parse the requested key out of it into TOut.
// This is used when you wish to add configuration keys that JanOS isn't aware of, but you'd
// still like to reference them from the same atlas file.
func Parse[TOut any](key string) TOut {
	refresh()

	gate.Lock()
	defer gate.Unlock()
	var zero TOut

	// 0. pull the requested key out, if possible

	value, exists := keys[key]
	if !exists {
		return zero
	}

	if typedValue, ok := value.(TOut); ok {
		return typedValue
	}

	// 1. if the found value isn't type-assertable, unmarshall it instead

	valueBytes, err := json.Marshal(value)
	if err != nil {
		return zero
	}

	var result TOut
	if err = json.Unmarshal(valueBytes, &result); err != nil {
		return zero
	}

	return result
}
