package JanOS

import (
	"fmt"
	"github.com/ignite-laboratories/JanOS/util"
	"log"
	"time"
)

// TODO: Add in log levels so we don't clutter the end user's console if it is not desired
// (We heavily use these logs to help us during development)

type logManager struct {
	History []LogEntry
}

// LogEntry represents a log entry at a single moment in time
type LogEntry struct {
	Time    time.Time
	Source  util.Named
	Message string
}

func newLogManager() *logManager {
	return &logManager{
		History: make([]LogEntry, 0),
	}
}

// AddEntry adds the provided log data to the historical record.
func (mgr *logManager) AddEntry(named util.Named, str string) {
	now := time.Now()
	mgr.History = append(mgr.History, LogEntry{
		Time:    now,
		Source:  named,
		Message: str,
	})
}

// Printf calls log.Print and captures the event on the Universe.os.Logging.History
func (mgr *logManager) Printf(source util.Named, format string, v ...any) {
	mgr.Println(source, fmt.Sprintf(format, v...))
}

// Println calls log.Println and captures the event on the Universe.Logging.History
func (mgr *logManager) Println(source util.Named, str string) {
	mgr.AddEntry(source, str)
	log.Printf("[%s] %s\n", source.GetName(), str)
}
