package bounded

type INumeric interface {
	ValueAsAny() any
	MinimumAsAny() any
	MaximumAsAny() any
	Range() uint64

	SetUsingAny(any) error
	SetBoundariesUsingAny(minimum, maximum any) error
}
