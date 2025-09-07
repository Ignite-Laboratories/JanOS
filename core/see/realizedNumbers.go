package see

// RealizedNumbers
/*
# " 𝑅𝑒𝑎𝑙𝑖𝑧𝑒𝑑 𝑁𝑢𝑚𝑏𝑒𝑟𝑠 "

A realized number is a "real number" driven by a neural activation path.  It holds the value as a calculated result
of dynamic precision and base, allowing it to identify and handle conditions such as 'irrationality' or 'periodicity'. The
reason it's a "realized" type, and not a "real number" type, is that this kind of number must be OBSERVED - and two
observers are not guaranteed to receive the same result.

Let me explain - in a static calculator, absolutely NONE of that is critical - but in a neural architecture, conditions
may need to be met for a calling entity to reveal the result.  Without an advanced numeric type like this, numbers must be
statically passed around - meaning they can no longer dynamically change with their environment.  A realized number
achieves this by implementing the see.Neuron interface, with the standard 𝑅𝑒𝑣𝑒𝑎𝑙() function driving an internalized
𝑟𝑒𝑣𝑒𝑙𝑎𝑡𝑖𝑜𝑛() function which updates the currently observed value.

	type Realized struct {
	  gate sync.Mutex

	  transcendental transcendental.Number

	  irrational    bool
	  negative      bool
	  whole         *Natural
	  fractional    *Natural
	  periodicWidth uint

	  revelation func(Realization) Realization
	  potential  func() bool

	  precision uint
	  base      uint16
	}

While the above structure may change slightly in code, I've circled around the concept enough to concretely say that this
will abstractly remain the same.  In pseudocode, on 𝐼𝑚𝑝𝑢𝑙𝑠𝑒() or 𝑅𝑒𝑣𝑒𝑎𝑙() this is what generally happens:

	if realized.potential() {
	  realized.gate.Lock()
	  defer realized.gate.Unlock()

	  realized = realized.revelation(copy(realized)) // self-realization =)
	}

The revelation function is provided a Realization - a copy of the Realized structure used for integration and
differentiation - and must return its calculated result.

The next thing you'll notice is that all fields are intentionally private in a realized number - this is because modification
of the values can have a direct influence upon other components: changing the base directly changes the periodic component,
for instance.  Thus, all access is provided through method calls available off the realized number type.
*/
type RealizedNumbers byte
