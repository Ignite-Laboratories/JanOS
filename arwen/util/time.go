package util

import (
	"fmt"
)

type TimeRecord struct {
	Step  int
	Value any
}

func (tr TimeRecord) String() string {
	return fmt.Sprintf("[%d=>%v]", tr.Step, tr.Value)
}

func NewTimeRecord(index int, value any) *TimeRecord {
	return &TimeRecord{
		Step:  index,
		Value: value,
	}
}
