package tiny

// A Realization is used to communicate a real number.  see.RealizedNumbers
//
// NOTE: If you'd like to emit side effects from your calculation, use the Artifact field.
type Realization struct {
	// Identity represents the symbolic identifier used to represent this value.
	Identity string

	// Irrational indicates if the value was observed to be irrational.
	Irrational bool

	// Negative indicates if the value is negative.
	Negative bool

	// Whole represents the digits constituting the whole part of this number.
	Whole []byte

	// Fractional represents the digits constituting the fractional part of this number.
	Fractional []byte

	// Periodic represents the digits constituting the periodic part of this number.
	Periodic []byte

	// Artifact represents any side effects emitted by the revelation function's activation.
	Artifact any

	// Base is a value from 2-256 indicating what base range the Whole/Fractional/Periodic digits are limited to.
	Base uint16

	// Precision indicates what precision the digits were calculated to.
	Precision uint

	// Revelation is a pointer to the RevelationFn that derived this Realization.
	Revelation RevelationFn

	// Identities is a map of the symbolic identities used to derive the Realization.
	//
	// NOTE: This can be used to cross-reference against the realization's Identity.
	Identities map[string]Realization
}

func (r Realization) String() string {

}
