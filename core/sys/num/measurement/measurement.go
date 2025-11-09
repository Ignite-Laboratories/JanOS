package measurement

import "git.ignitelabs.net/janos/core/sys/num"

func Random(bitWidth uint) num.Measurement {

}

//Outline:
//
//	0. Add in an ability to randomly generate measurements
//	1. Add in a way to diminish a measurement (a mutation - keeping the pattern as a separate artifact)
//	2. Build a process to midpoint with the following rules:
//		0. If the result yields 2 or more bits, place a terminus at the beginning
//		1. If the result takes the entire width of the index
//
//
//Bit Attrition:
//	Recursively midpointing an index to yield an average of over 1 bit of reduction.  At each step, a single bit
//is appended to the data indicating if it did (1) shrink by more than 1 bit, or didn't (0).  When you reach your
//final distilled width, you simply tell the system when you have reached either your target bit width or the number
//of steps you reduced for.  Then, you use the least significant bit to determine the next larger index size relative
//to the last.
//
//	There's a practical minimum to this - my guess is 64 bits - where if the target (minus the reduction bit) has
//fallen below that bit width then you'd possibly consider the data differently.
