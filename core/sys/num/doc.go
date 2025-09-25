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

Transcendentals are 'cached constants' - this means that they hold an anonymous method which makes them irrational,
but they are constantly referenceable from the num package as: num.Pi and num.E
*/
package num
