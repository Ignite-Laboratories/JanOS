package JanOS

import (
	"time"
)

type logManager struct {
	History []LogEntry
}

// LogEntry represents a log entry at a single moment in time
type LogEntry struct {
	Time    time.Time
	Named   named
	Message string
}

func newLogManager() *logManager {
	return &logManager{
		History: make([]LogEntry, 0),
	}
}

// GetNamedValue returns the assigned name to this instance.
func (mgr *logManager) GetName() string {
	return "Log"
}

// AddEntry adds the provided log data to the historical record.
func (mgr *logManager) AddEntry(named named, str string) {
	now := time.Now()
	mgr.History = append(mgr.History, LogEntry{
		Time:    now,
		Named:   named,
		Message: str,
	})
}
