package num

// A Realization is a mutable structure used by a Realized revelation to safely communicate a message.  see.RealizedNumbers
type Realization struct {
	Irrational bool
	Negative   bool
	Whole      []byte
	Fractional []byte
	Periodic   []byte
}

func (r Realization) String() string {

}
