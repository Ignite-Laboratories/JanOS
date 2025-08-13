package internal

import (
	"github.com/ignite-laboratories/core/std"
)

func NewPattern[T any](x std.Emit[T], data ...T) std.Pattern[T] {
	return std.Pattern[T]{
		Data: data,
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
	}
}

func NewPattern2D[T any](x std.Emit[T], y std.Emit[[]T], data ...[]T) std.Pattern2D[T] {
	return std.Pattern2D[T]{
		Data: data,
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
	}
}

func NewPattern3D[T any](x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], data ...[][]T) std.Pattern3D[T] {
	return std.Pattern3D[T]{
		Data: data,
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
		Z: std.Axis[[][]T]{
			Emit:   z,
			Cursor: std.NewCursorDefault(),
		},
	}
}

func NewPattern4D[T any](x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T], data ...[][][]T) std.Pattern4D[T] {
	return std.Pattern4D[T]{
		Data: data,
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
		Z: std.Axis[[][]T]{
			Emit:   z,
			Cursor: std.NewCursorDefault(),
		},
		W: std.Axis[[][][]T]{
			Emit:   w,
			Cursor: std.NewCursorDefault(),
		},
	}
}

func NewPattern5D[T any](x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T], a std.Emit[[][][][]T], data ...[][][][]T) std.Pattern5D[T] {
	return std.Pattern5D[T]{
		Data: data,
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
		Z: std.Axis[[][]T]{
			Emit:   z,
			Cursor: std.NewCursorDefault(),
		},
		W: std.Axis[[][][]T]{
			Emit:   w,
			Cursor: std.NewCursorDefault(),
		},
		A: std.Axis[[][][][]T]{
			Emit:   a,
			Cursor: std.NewCursorDefault(),
		},
	}
}

func NewPattern6D[T any](x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T], a std.Emit[[][][][]T], b std.Emit[[][][][][]T], data ...[][][][][]T) std.Pattern6D[T] {
	return std.Pattern6D[T]{
		Data: data,
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
		Z: std.Axis[[][]T]{
			Emit:   z,
			Cursor: std.NewCursorDefault(),
		},
		W: std.Axis[[][][]T]{
			Emit:   w,
			Cursor: std.NewCursorDefault(),
		},
		A: std.Axis[[][][][]T]{
			Emit:   a,
			Cursor: std.NewCursorDefault(),
		},
		B: std.Axis[[][][][][]T]{
			Emit:   b,
			Cursor: std.NewCursorDefault(),
		},
	}
}

func NewPattern7D[T any](x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T], a std.Emit[[][][][]T], b std.Emit[[][][][][]T], c std.Emit[[][][][][][]T], data ...[][][][][][]T) std.Pattern7D[T] {
	return std.Pattern7D[T]{
		Data: data,
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
		Z: std.Axis[[][]T]{
			Emit:   z,
			Cursor: std.NewCursorDefault(),
		},
		W: std.Axis[[][][]T]{
			Emit:   w,
			Cursor: std.NewCursorDefault(),
		},
		A: std.Axis[[][][][]T]{
			Emit:   a,
			Cursor: std.NewCursorDefault(),
		},
		B: std.Axis[[][][][][]T]{
			Emit:   b,
			Cursor: std.NewCursorDefault(),
		},
		C: std.Axis[[][][][][][]T]{
			Emit:   c,
			Cursor: std.NewCursorDefault(),
		},
	}
}
