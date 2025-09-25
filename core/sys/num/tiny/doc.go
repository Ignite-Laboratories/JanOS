// Package tiny
/*
𝑡𝑖𝑛𝑦 is an infinitely precise calculator (at least, as precise as your computer can afford)

It does this by performing arithmetic naïvely - as a child would.  This is by design as the underlying goal
of 𝑡𝑖𝑛𝑦 is to demystify mathematical processes for our future children.  I aim to produce a calculator that
provides the ability to "show its work" and act as a kind of "living proof engine".  This should also act
as the foundation for performing advanced std.Vector arithmetic using the num.Realizedand num.Natural types.

𝑡𝑖𝑛𝑦 comes with a few "limitations" - most importantly, bases.  Every placeholder digit is a single byte value,
allowing it to hold up to base₂₅₆.  Fundamentally, that's okay, but how do we -represent- base₂₅₆?  Or any
base above base₁₆ [0-F]?  By doubling up the value of each placeholder as a hex value.  This gives us the
closed interval of [00,FF] for string representations of a symmetrically sized byte value [0,255].  As base₂₅₆
requires 256 identifiers, that's our fundamental cap without doubling the size of all placeholders using an int16.
Since we track each placeholder as a byte value, a space character is placed between all placeholders for any
string representation of a num.Realizedabove base₁₆ - whereas the underlying placeholder byte slice naturally
"separates" each for the computer.

A num.Realizedis composed of three components - the whole num.Natural part, the fractional num.Natural part, and
the periodic region.  The periodic region denotes the width of the end fractional part which repeats infinitely.
*/
package tiny
