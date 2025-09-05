// Package num
/*

The `num` package is a bridge between numeric Go types and 𝑡𝑖𝑛𝑦.

This provides access to the Numeric type, which provides a way to create "boundable" generically typed numbers.

`num` also provides several interfaces for describing numeric-compatible types:
Primitive, Advanced, Signed, Integer, Float, Complex

The Primitive interface retains all mathematical operators, whereas the Advanced interface provides access to
the 𝑡𝑖𝑛𝑦 types - which can be explored through tiny.Realized

There are many convenience methods in this package:

Random number generation -
Random, RandomWithinRange, RandomSet, RandomSetWithinRange

Anonymous procedures -
ToString, TypeAssert, Smallest, Largest, Compare, IsNumeric, IsPrimitive, IsSigned, IsFloat, IsComplex, IsInteger, IsInf, and IsNaN

Minimum/Maximums -
MinValue[Primitive], MaxValue[Primitive]

---------
Transcendentals

Transcendentals are commonly used values which need not be calculated "on the fly" - however, 𝑡𝑖𝑛𝑦 is an infinitely
precise calculator!  That means these constants HAVE to be calculated "on the fly!"  Otherwise, whatever value we
calculate will get rounded off to a fixed placeholder width - thus, the transcendentals are 'cached constants' -
meaning they are calculated ONCE for each requested placeholder width "on demand."
*/
package num
