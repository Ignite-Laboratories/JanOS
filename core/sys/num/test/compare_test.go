package test

import (
	"core/sys/num"
	"math"
	"testing"
)

func Test_Compare(t *testing.T) {
	tests := []struct {
		name        string
		a, b        any
		want        int
		shouldPanic bool
	}{
		{"NaN vs number", math.NaN(), 0.0, -1, false},
		{"number vs NaN", 0.0, math.NaN(), 1, false},
		{"NaN vs NaN", math.NaN(), math.NaN(), 1, true},
		{"+Inf vs +Inf", math.Inf(1), math.Inf(1), 0, false},
		{"-Inf vs -Inf", math.Inf(-1), math.Inf(-1), 0, false},
		{"-Inf vs finite", math.Inf(-1), 0.0, -1, false},
		{"+Inf vs finite", math.Inf(1), 0.0, 1, false},
		{"finite vs -Inf", 0.0, math.Inf(-1), 1, false},
		{"finite vs +Inf", 0.0, math.Inf(1), -1, false},
		{"-0 vs +0", math.Copysign(0.0, -1), 0.0, 0, false},
		{"-0.0 vs 0.000", -0.0, 0.000, 0, false},
		{"-0 vs 5", -0.0, 5.0, -1, false},
		{"-1 vs -0.9", -1.0, -0.9, -1, false},
		{"-0.01 vs -0.1", -0.01, -0.1, 1, false},
		{"5.0505 vs 5.55", 5.0505, 5.55, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					if tt.shouldPanic {
						t.Errorf("The code did not panic")
					}
				}
			}()
			result := num.Compare(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("Compare(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.want)
			}
		})
	}
}
