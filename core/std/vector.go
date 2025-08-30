package std

import "core/sys/num/bounded"

// Vector represents any "𝑙𝑒𝑡𝑡𝑒𝑟 𝑣𝑒𝑐𝑡𝑜𝑟" type, and all generated letter vectors create these functions at creation. See letters.Doc
//
// This provides a type-agnostic way of retrieving and setting vector information at runtime.  These
// will intentionally panic if provided erroneous parameters, such as trying to 'Set' with a mismatched type.
//
// All vectors are components of bounded.Number values - vector operations pay no mind to the -actual- types involved,
// so long as their primitive arithmetic operations are available.  This means you will need to perform type casting
// to -your- bounded.Numeric type, which should be available to you.
//
// NOTE: By design, all vector components are organized by their self-described order.  For instance, an XYZW type's
// components are logically ordered as X[0], Y[1], Z[2], W[3] - while a UV's is U[0], V[1].
//
// NOTE: Name retrieval is case-insensitive - "YCb" and "yCb" both match YCbCr's "Cb" component.
type Vector interface {
	// ComponentLen returns the number of components in the vector.
	ComponentLen() uint

	// Components returns a slice of the bounded.Number components in the vector.
	Components() []bounded.INumeric

	// Component gets the component value at the provided index.
	//
	// NOTE: By design, all vector components are organized by their self-described order.  For instance, an XYZW type's
	// components are logically ordered as X[0], Y[1], Z[2], W[3] - while a UV's is U[0], V[1].
	Component(uint) (bounded.INumeric, error)

	// ComponentByName gets the named component's value.
	//
	// NOTE: Name retrieval is case-insensitive - "YCb" and "yCb" both match YCbCr's "Cb" component.
	ComponentByName(string) (bounded.INumeric, error)

	// SetComponents treats the input slice as a 1:1 mapping with the underlying components of the appropriate types before setting their values.
	SetComponents([]any) error

	// SetComponent sets the component value at the provided index.
	//
	// NOTE: By design, all vector components are organized by their self-described order.  For instance, an XYZW type's
	// components are logically ordered as X[0], Y[1], Z[2], W[3] - while a UV's is U[0], V[1].
	SetComponent(uint, any) error

	// SetComponentByName sets the named component's value.
	//
	// NOTE: Name retrieval is case-insensitive - "YCb" and "yCb" both match YCbCr's "Cb" component.
	SetComponentByName(string, any) error

	String() string
}
