package dimension

import (
	"github.com/ignite-laboratories/core/enum/dimension/aaxis"
	"github.com/ignite-laboratories/core/enum/dimension/baxis"
	"github.com/ignite-laboratories/core/enum/dimension/caxis"
	"github.com/ignite-laboratories/core/enum/dimension/waxis"
	"github.com/ignite-laboratories/core/enum/dimension/xaxis"
	"github.com/ignite-laboratories/core/enum/dimension/yaxis"
	"github.com/ignite-laboratories/core/enum/dimension/zaxis"
	"github.com/ignite-laboratories/core/enum/direction/ordinal"
	"reflect"
)

// Any is a way of referencing any normalized axial dimension.  Every axis provides a 'Negative', 'Static', and 'Positive'
// type - representing the normalized points of '-1', '0', and '1' respectively.  By providing these types, you may now
// constrain your functions to any arrangement of these points for your purposes.  If you wish to combine the axes together
// you may do so through the multidimensional 'any' types.
//
// See Any, Any2D, Any3D, Any4D, Any5D, Any6D, and Any7D
type Any interface {
	XAxis | YAxis | ZAxis | WAxis | AAxis | BAxis | CAxis
}

type XAxis interface {
	xaxis.Negative | xaxis.Static | xaxis.Positive
}

type YAxis interface {
	yaxis.Negative | yaxis.Static | yaxis.Positive
}

type ZAxis interface {
	zaxis.Negative | zaxis.Static | zaxis.Positive
}

type WAxis interface {
	waxis.Negative | waxis.Static | waxis.Positive
}

type AAxis interface {
	aaxis.Negative | aaxis.Static | aaxis.Positive
}

type BAxis interface {
	baxis.Negative | baxis.Static | baxis.Positive
}

type CAxis interface {
	caxis.Negative | caxis.Static | caxis.Positive
}

// GetDepth returns the dimensional depth of the provided []T.
func GetDepth[T any](v []T) int {
	t := reflect.TypeOf(v)
	d := 0
	for t.Kind() == reflect.Slice {
		d++
		t = t.Elem()
	}
	return d
}

// AsOrdinal converts the provided axial type to a concrete ordinal.Direction, which is a byte valued -1, 0, or 1.
//
// For example -
//
//	switch dimension.AsOrdinal[TDim]() {
//	case -1:
//	case 0:
//	case 1:
//	default:
//	}
func AsOrdinal[T Any]() ordinal.Direction {
	switch any(T(0)).(type) {
	case xaxis.Negative, yaxis.Negative, zaxis.Negative, waxis.Negative, aaxis.Negative, baxis.Negative, caxis.Negative:
		return ordinal.Negative
	case xaxis.Static, yaxis.Static, zaxis.Static, waxis.Static, aaxis.Static, baxis.Static, caxis.Static:
		return ordinal.Static
	case xaxis.Positive, yaxis.Positive, zaxis.Positive, waxis.Positive, aaxis.Positive, baxis.Positive, caxis.Positive:
		return ordinal.Positive
	default:
		panic("invalid type")
	}
}

// Any2D represents any set of two dimensionally normal points - a "plane".
//
// Multidimensional slices represent the 'first' dimension as the innermost slice.  Thus, these are the named axial dimensions
// and what JanOS considers the structure to be called.
//
//	X - Line          - []T
//	Y - Plane         - [][]T
//	Z - Volume        - [][][]T
//	W - Tesseract     - [][][][]T
//	A - Awareness     - [][][][][]T
//	B - Consciousness - [][][][][][]T
//	C - Universe      - [][][][][][][]T
//
// See Any, Any2D, Any3D, Any4D, Any5D, Any6D, and Any7D
type Any2D[T1 Any, T2 Any] struct {
	First  T1
	Second T2
}

func New2D[T1 Any, T2 Any](primary T1, secondary T2) Any2D[T1, T2] {
	return Any2D[T1, T2]{primary, secondary}
}

// Any3D represents any set of three dimensionally normal points - a "volume".
//
// Multidimensional slices represent the 'first' dimension as the innermost slice.  Thus, these are the named axial dimensions
// and what JanOS considers the structure to be called.
//
//	X - Line          - []T
//	Y - Plane         - [][]T
//	Z - Volume        - [][][]T
//	W - Tesseract     - [][][][]T
//	A - Awareness     - [][][][][]T
//	B - Consciousness - [][][][][][]T
//	C - Universe      - [][][][][][][]T
//
// See Any, Any2D, Any3D, Any4D, Any5D, Any6D, and Any7D
type Any3D[T1 Any, T2 Any, T3 Any] struct {
	First  T1
	Second T2
	Third  T3
}

func New3D[T1 Any, T2 Any, T3 Any](primary T1, secondary T2, third T3) Any3D[T1, T2, T3] {
	return Any3D[T1, T2, T3]{primary, secondary, third}
}

// Any4D represents any set of four dimensionally normal points - a "tesseract".
//
// Multidimensional slices represent the 'first' dimension as the innermost slice.  Thus, these are the named axial dimensions
// and what JanOS considers the structure to be called.
//
//	X - Line          - []T
//	Y - Plane         - [][]T
//	Z - Volume        - [][][]T
//	W - Tesseract     - [][][][]T
//	A - Awareness     - [][][][][]T
//	B - Consciousness - [][][][][][]T
//	C - Universe      - [][][][][][][]T
//
// See Any, Any2D, Any3D, Any4D, Any5D, Any6D, and Any7D
type Any4D[T1 Any, T2 Any, T3 Any, T4 Any] struct {
	First  T1
	Second T2
	Third  T3
	Fourth T4
}

func New4D[T1 Any, T2 Any, T3 Any, T4 Any](primary T1, secondary T2, third T3, fourth T4) Any4D[T1, T2, T3, T4] {
	return Any4D[T1, T2, T3, T4]{primary, secondary, third, fourth}
}

// Any5D represents any set of five dimensionally normal points.
//
// Multidimensional slices represent the 'first' dimension as the innermost slice.  Thus, these are the named axial dimensions
// and what JanOS considers the structure to be called.
//
//	X - Line          - []T
//	Y - Plane         - [][]T
//	Z - Volume        - [][][]T
//	W - Tesseract     - [][][][]T
//	A - Awareness     - [][][][][]T
//	B - Consciousness - [][][][][][]T
//	C - Universe      - [][][][][][][]T
//
// See Any, Any2D, Any3D, Any4D, Any5D, Any6D, and Any7D
type Any5D[T1 Any, T2 Any, T3 Any, T4 Any, T5 Any] struct {
	First  T1
	Second T2
	Third  T3
	Fourth T4
	Fifth  T5
}

func New5D[T1 Any, T2 Any, T3 Any, T4 Any, T5 Any](primary T1, secondary T2, third T3, fourth T4, fifth T5) Any5D[T1, T2, T3, T4, T5] {
	return Any5D[T1, T2, T3, T4, T5]{primary, secondary, third, fourth, fifth}
}

// Any6D represents any set of six dimensionally normal points.
//
// Multidimensional slices represent the 'first' dimension as the innermost slice.  Thus, these are the named axial dimensions
// and what JanOS considers the structure to be called.
//
//	X - Line          - []T
//	Y - Plane         - [][]T
//	Z - Volume        - [][][]T
//	W - Tesseract     - [][][][]T
//	A - Awareness     - [][][][][]T
//	B - Consciousness - [][][][][][]T
//	C - Universe      - [][][][][][][]T
//
// See Any, Any2D, Any3D, Any4D, Any5D, Any6D, and Any7D
type Any6D[T1 Any, T2 Any, T3 Any, T4 Any, T5 Any, T6 Any] struct {
	First  T1
	Second T2
	Third  T3
	Fourth T4
	Fifth  T5
	Sixth  T6
}

func New6D[T1 Any, T2 Any, T3 Any, T4 Any, T5 Any, T6 Any](primary T1, secondary T2, third T3, fourth T4, fifth T5, sixth T6) Any6D[T1, T2, T3, T4, T5, T6] {
	return Any6D[T1, T2, T3, T4, T5, T6]{primary, secondary, third, fourth, fifth, sixth}
}

// Any7D represents any set of seven dimensionally normal points.
//
// Multidimensional slices represent the 'first' dimension as the innermost slice.  Thus, these are the named axial dimensions
// and what JanOS considers the structure to be called.
//
//	X - Line          - []T
//	Y - Plane         - [][]T
//	Z - Volume        - [][][]T
//	W - Tesseract     - [][][][]T
//	A - Awareness     - [][][][][]T
//	B - Consciousness - [][][][][][]T
//	C - Universe      - [][][][][][][]T
//
// See Any, Any2D, Any3D, Any4D, Any5D, Any6D, and Any7D
type Any7D[T1 Any, T2 Any, T3 Any, T4 Any, T5 Any, T6 Any, T7 Any] struct {
	First   T1
	Second  T2
	Third   T3
	Fourth  T4
	Fifth   T5
	Sixth   T6
	Seventh T7
}

func New7D[T1 Any, T2 Any, T3 Any, T4 Any, T5 Any, T6 Any, T7 Any](primary T1, secondary T2, third T3, fourth T4, fifth T5, sixth T6, seventh T7) Any7D[T1, T2, T3, T4, T5, T6, T7] {
	return Any7D[T1, T2, T3, T4, T5, T6, T7]{primary, secondary, third, fourth, fifth, sixth, seventh}
}
