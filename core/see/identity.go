package see

// Identity
/*
# " 𝐼𝑑𝑒𝑛𝑡𝑖𝑡𝑦 "

While the concept of identity within JanOS is complex and layered, it can be distilled down to two unique conditions:

- Entity Identification

- Numeric Identification

Entities, and their associated naming system, are covered in a different document.  This document covers the numeric
identity system used by 𝑡𝑖𝑛𝑦.

# Self-Realization

The concept of a realized number is covered in see.RealizedNumbers - however, the concept of identification is not
covered there, as it's far broader.

Essentially - how do you identify if something is "irrational"?

You typically need to follow a symbolic trail that proves whether it's an irrational value!  The realized type
already provides a way to dynamically create an irrational value, but even IT must follow a symbolic trail to
prove irrationality.  Thus, all realized numbers self-identify their symbolic trail during realization - it's
simply the ONLY way this kind of architecture could be maintainable in the long run!

They must intelligently identify themselves, though!  Thus, what does it even mean to 'identify' oneself?
Our actions define our identifiable selves through a sea of noise - just as the results of a revelation do.

Luckily, we're in the land of -numeric- identification, at least for the moment - we can follow a standardized
convention that allows the realized identity to be machine-parseable.  A realized number like π is easy to
"identify" from its Greek "identity" - but an arbitrary numeric value doesn't have a fancy letter character
to represent itself, it must shine as it exists!

	 "π"    ← Pi
	 "1"    ← The loneliest number
	"42.00" ← The number '42'

Literally, through the concept of see.PrintingNumbers any numeric value can self-identify!  But that doesn't
-quite- get us a symbolic trail.  Instead, the entire formula must be output in the resulting realization's
identity field:

	"(π + 1)" ← An addition identity

The above is intuitively parseable - it contains two operand identities linked by an operator.  This becomes
the realization's 'identity', and it's separate from the realized value, which expands out π to its constituent
digits.  By associating the identity with the calculated digits, we can retain proof of why we deemed the value
to be "irrational" - even though we HAD to eventually cut off the digits at some point.

Since the realized value is generated on-demand, any change in precision or base could yield different rationality -
thus the resulting realization always comes from activating the revelation function, guaranteeing an always provable
result.

It doesn't JUST provide the identities, however - the realization ALSO provides a map associating the printed
identities with the underlying realized number.  This gives us two features:

1. The realization self-describes how to regenerate each symbol in the identity

2. Revelations can cross-reference identities between realized results


*/
type Identity byte
