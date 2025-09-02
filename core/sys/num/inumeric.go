package num

// INumeric allows Numeric types to be accessed anonymously by providing 'any' typed functions.
type INumeric interface {
	ValueAsAny() any
	MinimumAsAny() any
	MaximumAsAny() any
	Range() uint64

	SetUsingAny(any) Breach
	SetBoundariesUsingAny(minimum, maximum any) Breach
}
