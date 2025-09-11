package see

// PrintingNumbers
/*
# " 𝑃𝑟𝑖𝑛𝑡𝑖𝑛𝑔 𝑁𝑢𝑚𝑏𝑒𝑟𝑠 "

In the context of 𝑡𝑖𝑛𝑦, printing refers to either the output of a number OR the output of a binary measurement
and/or phrase.  This document covers printing numbers.  For binary, see.PrintingBinary

# Printing Operations

In short, there are two kinds of numbers in 𝑡𝑖𝑛𝑦 - the Natural and the Realized.  Both of these offer three operations:

	𝑆𝑡𝑟𝑖𝑛𝑔() - Outputs the data in identity form at the observed base
	𝑃𝑟𝑖𝑛𝑡()  - Outputs the data in human-legible form as any base
	𝑀𝑎𝑡𝑟𝑖𝑥() - Outputs the data in well-aligned and machine-legible form for matrix operations

For example, let's look at the three variants of outputting four unique types:

	 π - An identified irrational
	√3 - An unidentified irrational calculation
	 𝑥 - A periodic [77.7777...]
	 𝑦 - A static number [1]

First, we have the general 𝑆𝑡𝑟𝑖𝑛𝑔() operation:

	  ⬐ "Print the operand in the most concise form, please"
	 "π"             ←  π.String()
	"~1.7320508"     ← √3.String() NOTE: Irrationals string to atlas.PrecisionMinimum - default 7
	"77.‾7"          ←  𝑥.String()
	 "1"             ←  𝑦.String()

Second, we have the 𝑃𝑟𝑖𝑛𝑡() operation:

	  ⬐ "Print the operand to 11 fractional places (and optionally in base₁₀), please"
	"~3.14159265359" ← π.Print(11, [10])
	"~1.73205080757" ← √3.Print(11, [10])
	"77.77777777778" ← 𝑥.Print(11, [10])
	 "1.00000000000" ← 𝑦.Print(11, [10])

	NOTE: If you'd like to print whatever width has already been calculated, use a width of '-1'

Lastly, we have the 𝑀𝑎𝑡𝑟𝑖𝑥() operation:

	  ⬐ "Print the operands √3, 𝑥, and 𝑦 as a matrix aligned against π to 11 places in base₁₀, please"
	"01.73205080757"
	"22.22222222222"
	"01.00000000000"
	"03.14159265359" ← π.Matrix(11, 10, √3, 𝑥, 𝑦)

	NOTE: Using a width of -1 aligns to the width of the widest fractional operand's current calculation

Print operations explicitly allow you to render a number to a specified precision and/or base.  Matrix operations
do so while aligning the operands relatively to one another.  There is an important, but subtle, feature of matrix
operations: they are aligned -against- a specific operand!  When doing so, the anchoring operand is referred to as
the matrix's 'solution' - as its output in the final row.  When performing arithmetic, 𝑡𝑖𝑛𝑦 uses this feature heavily
to "pre-allocate" a solution row against the provided operands.

# Bases

The output text treats every placeholder as a hexadecimal number.  This is because 𝑡𝑖𝑛𝑦 stores all numbers in
base₂ (for multiple reasons - but, mostly, to make the binary synthesis infinitely easier).  In doing so, the very
act of naïve calculation against a matrix forces each bit to occupy a full byte of memory in Go (at least, during
calculation.) Rather than fight this, we use it to our advantage!  A byte affords us a 'placeholder' from base₂
to base₂₅₆ for all operations using a single type.  The best feature of a byte is that we can output its value
in hexadecimal form for -all- digit values, through a standardized convention:

  0. For base₁₆ and below, every digit is a single-character hexadecimal value.
  1. For base₁₇ and above, every digit is a two-character hexadecimal value.
  2. For base₁₇ and above, every single position, including [~-.‾], is spaced with whitespace.

# String Operations In Depth

Let's take a look at the simplest 𝑡𝑖𝑛𝑦 type - the Natural.  As its name implies, the natural describes a 'natural
number' - or any positive whole countable number, INCLUDING zero. [ Briefly - I'd like to note that the higher order bases
all rely upon the value of 'zero' - otherwise, base₁ couldn't even exist, let alone any higher-order base! How else
could you represent binary!?  Arggg ] I digress - apologies, my convictions get the better of me from time to time -
back to printing numbers!

	" A natural 42 in several bases "

	"101010" ← base₂
	    "52" ← base₈
	    "42" ← base₁₀
	    "36" ← base₁₂
	    "2A" ← base₁₆
	 "02 08" ← base₁₇
	 "01 14" ← base₂₂

[ I should note, 42 coincidentally has a beautiful "candy-striped" binary pattern =) ]

Next, from the natural, we have the "realized number."  A realized number, in essence, is a "Real Number" consisting
of four parts:

	0. The sign
	1. The whole part
	2. The fractional part
	3. The periodic part

Realized numbers print as you would expect → "whole.fractional‾periodic" (with an optional leading minus sign)

However, a "realized" number earns its name because it can -dynamically- recognize an irrational or periodic condition,
which you can read about in see.RealizedNumbers.  For printing, that means we have several unique conditions to print:

	 "123"        ← An integer value
	"-123"        ← An negative integer value

	 "123.45"     ← A floating point value
	"-123.45"     ← A negative floating point value

	 "123.‾4"     ← A periodic value, broken with an "overscore"
	"-123.‾4"     ← A negative periodic value, broken with an "overscore"

	 "123.‾456"   ← A "wide" periodic value
	"-123.‾456"   ← A "wide" negative periodic value

	 "123.45‾678" ← An mixed-repeating periodic value
	"-123.45‾678" ← A negative mixed-repeating periodic value

	   "π"        ← The identity of Pi
	   "ℯ"        ← The identity of Euler's number

	 "~1.7320508" ← An irrational value to atlas.PrecisionMinimum digits (default 7) [√3]
	"~-1.7320508" ← A negative irrational value to atlas.PrecisionMinimum digits (default 7) [-√3]

Several key points should be identified:

0. Irrationals cannot be realized to any width, thus they are dynamically generated and eventually rounded - by default, to atlas.Precision

1. Irrationals are identified with a tilde character [~] and rounded to atlas.PrecisionMinimum fractional places (default 7) for concise output.

2. Periodic values are broken with an "overscore" character [‾], indicating the pattern to the right infinitely repeats.

3. When using String(), identities will output (if assigned)

To summarize - String() is used to print concise readable output and is always printed in the stored base.  For natural numbers,
the stored base is ALWAYS base₁₀ - whereas for realized numbers, you can explicitly set the base on the number before calling String().

# Print Operations

The next operation 𝑡𝑖𝑛𝑦 affords are 𝑃𝑟𝑖𝑛𝑡() operations.  These act just as the 𝑆𝑡𝑟𝑖𝑛𝑔() operations, with two exceptions:

0. You can print the result in any base you desire

1. Identities are not output

2. Precision width can be set (NOTE: only realized numbers have precision)

When printing in a different base, periodic values are repeated out to maximum precision before conversion. Here's what
"-42.‾54321" [base₁₀] would look like in several different bases:

	"-101010.1101010000110001" ← base₂
	    "-52.152061"           ← base₈
	    "-42.‾54321"           ← base₁₀
	    "-36.27529"            ← base₁₂
	    "-2A.D431"             ← base₁₆
	"- 02 08 . 0B 00 10 06"    ← base₁₇
	"- 01 14 . 05 02 05 03"    ← base₂₂

Note that only the base where periodicity is known will output the overscore character.  Identified irrationals, such
as π, are printed without their identity:

	"~3.1415927" ← π [base₁₀] to atlas.PrecisionMinimum digits (default 7)

Realized numbers store their currently calculated fractional precision, but during a print operation you may override
that and request a different precision width.  Printing will not cause the number to recalculate to that precision, but
instead will either round the number early or pad it with zeros to your desired width.  If you provide a width of '-1',
the resulting number will be printed to whatever precision is currently calculated.

# Matrix Operations

These operations allow quick alignment of operands for matrix calculations.  Essentially, if given three
operands of different widths, they must be padded on both sides with "0"s for alignment:

	"1234.56789"
	"0001.00000"
	"0000.20000"

𝑀𝑎𝑡𝑟𝑖𝑥() operations differ from 𝑆𝑡𝑟𝑖𝑛𝑔() and 𝑃𝑟𝑖𝑛𝑡() in that:

	0. The [~‾] characters are never emitted
	1. The first placeholder is ALWAYS the sign, printed as [+-] characters
	2. Irrationals and periodics are always rounded to the desired precision
	3. Operands are aligned against each other

To the last point - to create a matrix, you must start with one operand and create a matrix by applying
others against it.  This allows the operands themselves to define the overall width of the matrix, rather than causing
undesired truncation of the whole or fractional parts.

*/
type PrintingNumbers byte
