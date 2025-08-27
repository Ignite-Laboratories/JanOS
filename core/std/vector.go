package std

// Vector represents any letter vector type - all letter vectors create these functions on generation.
//
// This provides a type-agnostic way of retrieving and setting vector information at runtime.  These
// will intentionally panic if provided erroneous parameters, such as trying to 'Set' with a mismatched type.
//
// NOTE: By design, all vector components are organized by their self-described order.  For instance, an XYZW type's
// components are logically ordered as X:0, Y:1, Z:2, W:3 - while a UV's is U:0, V:1.
//
// NOTE: Name retrieval is case-insensitive - "YCb" and "yCb" both match YCbCr's "Cb" component.
type Vector interface {
	// SetComponent sets the component value at the provided index.
	//
	// NOTE: By design, all vector components are organized by their self-described order.  For instance, an XYZW type's
	// components are logically ordered as X:0, Y:1, Z:2, W:3 - while a UV's is U:0, V:1.
	SetComponent(uint, any)

	// GetComponent gets the component value at the provided index.
	//
	// NOTE: By design, all vector components are organized by their self-described order.  For instance, an XYZW type's
	// components are logically ordered as X:0, Y:1, Z:2, W:3 - while a UV's is U:0, V:1.
	GetComponent(uint) any

	// SetComponentByName sets the named component's value.
	//
	// NOTE: Name retrieval is case-insensitive - "YCb" and "yCb" both match YCbCr's "Cb" component.
	SetComponentByName(string, any)

	// GetComponentByName gets the named component's value.
	//
	// NOTE: Name retrieval is case-insensitive - "YCb" and "yCb" both match YCbCr's "Cb" component.
	GetComponentByName(string) any

	String() string
}
