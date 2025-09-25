package std

// A Vector is any type that can dynamically retrieve named components.
//
// NOTE: Component names are case invariant!
//
// See ComponentLen, Components, and Component
// Vector represents any "𝑙𝑒𝑡𝑡𝑒𝑟 𝑣𝑒𝑐𝑡𝑜𝑟" type, and all generated letter vectors are given these functions at creation. See letters.Doc
//
// Vectors, abstractly, are any type that can dynamically retrieve named components.  The underlying type of each component is
// unique to each vector, but
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
	// ComponentLen ( ) retrieves the number of components managed by this Vector.
	//
	// See ComponentLen, Components, and Component
	ComponentLen() uint

	// Components ( ...𝑛𝑎𝑚𝑒𝑠 ) returns multiple components by 𝑛𝑎𝑚𝑒.
	//
	// NOTE: Component names are case invariant!
	//
	// See ComponentLen, Components, and Component
	Components(names ...string) map[string]any

	// Component ( 𝑛𝑎𝑚𝑒𝑑 ) returns a single 𝑛𝑎𝑚𝑒𝑑 component.
	//
	// NOTE: Component names are case invariant!
	//
	// See ComponentLen, Components, and Component
	Component(named string) any

	// Set ( 𝑛𝑎𝑚𝑒𝑑, 𝑣𝑎𝑙𝑢𝑒 ) assigns the provided value to the 𝑛𝑎𝑚𝑒𝑑 component.
	Set(named string, value any) Vector

	// From ( 𝑎𝑛𝑜𝑛𝑦𝑚𝑜𝑢𝑠 ) assigns any of the matching named anonymous components from the provided map.
	From(Anonymous) Vector

	String() string
}
