package main

import (
	"core/std"
	"core/sys/num/tiny"
	"reflect"
	"strings"
)

func main() {
	a := XYZW{X: 1, Y: 2, Z: 3, W: 4.44}
	b := std.Anonymous{"X": 5}
	c := std.Anonymous{"Y": 4, "Z": 3}
	a.Add(b, c)
}

type XYZW struct {
	X, Y, Z, W float64
}

// Add adds the like-named components of the operands.  If you'd like to mix-and-match your components, you
// may create an std.Anonymous map - for example:
func (v *XYZW) Add(operands ...std.Vector) *XYZW {
	if len(operands) == 0 {
		return v
	}
	for _, operand := range operands {
		for name, value := range operand.Components() {
			switch strings.ToLower(name) {
			case "x":
				v.X += tiny.Add[float64](v.X, value.(float64))
			case "y":
				v.Y += tiny.Add[float64](v.Y, value.(float64))
			case "z":
				v.Z += tiny.Add[float64](v.Z, value.(float64))
			case "w":
				v.W += tiny.Add[float64](v.W, value.(float64))
			}
		}
	}
	return v
}

func (v *XYZW) Components(names ...string) []any {
	return Lookup(v, names...)
}

func (v *XYZW) Component(named string) any {
	return Lookup(v, named)[0]
}

func Lookup[T any](v T, names ...string) []any {
	out := make([]any, len(names))
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return out
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return out
	}

	for i, named := range names {
		fv := rv.FieldByName(named)
		if !fv.IsValid() || !fv.CanInterface() {
			continue
		}
		out[i] = fv.Interface()
	}
	return out
}
