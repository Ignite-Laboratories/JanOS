package JanOS

import (
	"time"
)

type logManager struct {
	History []LogEntry
}

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

func (mgr *logManager) GetName() string {
	return "Log"
}

func (mgr *logManager) AddEntry(named named, str string) {
	now := time.Now()
	mgr.History = append(mgr.History, LogEntry{
		Time:    now,
		Named:   named,
		Message: str,
	})
}
