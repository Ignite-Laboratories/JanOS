package see

// ActionPotentials
/*
# " 𝐴𝑐𝑡𝑖𝑜𝑛 ← 𝑃𝑜𝑡𝑒𝑛𝑡𝑖𝑎𝑙𝑠 "

An action potential is any action gated by a rudimentary potential function.  A potential function is an anonymous
function that simply returns a boolean to indicate if a condition has been met.  The sympathetic action isn't driven
BY the potential function - rather, by its hosting entity.  This is because the separation of an 𝑎𝑐𝑡𝑖𝑜𝑛 from its
𝑝𝑜𝑡𝑒𝑛𝑡𝑖𝑎𝑙 allows fine-tuning of the potential function until the desired moment of activation is reached WITHOUT
mutating the action.

	" 𝐴 𝑃𝑜𝑡𝑒𝑛𝑡𝑖𝑎𝑙 "
	func SomePotential() bool { ... }

	" 𝐴𝑛 𝐴𝑐𝑡𝑖𝑜𝑛 "
	func SomeAction() { ... }

NOTE: Many standardized temporal potential functions are available in the "𝑤ℎ𝑒𝑛" package (see when.Doc)

The key point is that these action functions are usually gated by potential closures over pointers!
For instance, let's take a look at a "frequency activation" function:

	func Frequency[T num.Primitive](hertz *T) func() bool {
	  last := time.Now()

	  return func() bool {
	    now := time.Now()

	    if now.Sub(last) > HertzToDuration(*hertz) {
	      last = now
	      return true
	    }
	    return false
	  }
	}

Here, the input parameter is given as a -pointer- to the "hertz" value meant to gate this potential function.
Using this standardized function, clusters of many neural activations can be gated by the same cyclic frequency
- simply because each knows HOW to talk the language of 'time' in a standardized activation sequence.
*/
type ActionPotentials byte
