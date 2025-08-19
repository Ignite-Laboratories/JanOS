// Package recorded provides access to the concept of -experiencing- data.  Recorded provides a comprehensive set of
// semantic ways to interpret information across time - also known as -describing- experiences.  From an abstract
// perspective, we can distill experiences down to a set of identifiers that have been recorded by a specific entity.
//
// This creates a recorded timeline of experiential identifiers which should be queryable and integrated into the future
// decision-making process.  So, the perspective in mind when working with recorded is that data is experienced contextually
// across time - and, most of the time, should be seeded through the power of entropy.
//
// This presents us with a recursive stack of recorded data sets, each with a different perspective they provide to a larger
// algorithm.  The sets detailed below operate on the principle of 'yielding' a data point while recording whatever is
// contextually relevant.
//
//	0 - std.Data[T any]
//	Abstract: "Anonymous closure"-safe typed slices of data.
//
//	Details:
//
// The reason this came into creation is due to a basic limitation of an anonymous-driven
// architecture: slices inherently cannot be appended against from outside an anonymous
// function that closed on the variable.  In short, this is because the act of 'appending'
// re-assigns the header that actually points to the slice data.  By wrapping all slices
// of data in this type, you can safely pass std.Data structures into closures without any
// worry.  To achieve this, the abstract act of "appending to a slice" is intercepted by
// a method call and owned by the std.Data type throughout JanOS.  Thus, this belongs in
// the 'std' namespace.  In addition, all std.Data objects are a std.Entity - allowing them
// to be identified and named.
//
// This type also introduces the first recursive 'Yield(n) []T' method, which is driven by
// a provided 'Yield() T' function at the time of creation.  The idea is simple - YOU
// understand the concept of what you want to record!  Just tell us how you'd like to anonymously
// Yield a value, then let us handle the concept of managing "recording" the result.
// Each time Yield is called, the results are stored in the set before being returned to the
// caller.  This is also why I've chosen the package name 'recorded' over the more conventional
// 'stream' terminology =)
//
// NOTE: Advanced std.Data creation functions still exist within the 'recorded' namespace
//
//	 tl;dr - std.Data represents the abstract concept of 'recording observed entropy'
//
//		1 - recorded.Unique[T comparable]
//	 Abstract: "Give me something I haven't seen yet!"
//
//	 Details:
//
// This type came into creation because random number generation in a bounded space has a basic
// flaw: repetition!  For instance, randomly emitting numbers from the closed interval [5, 10]
// might yield '6 6 6' - which, for a system aimed to fabricate diverse experiences, represents
// STAGNATION!
//
// This type introduces the concept of 'periodic uniqueness' - or, exhausting a set of values
// uniquely through Yield() before repeating the next exhaustive round.  The details of HOW
// this mechanic works are defined on the type itself.  Most importantly, however, this type
// can be created in two ways: bounded, or seeded.  A bounded unique set simply yields random
// numbers within a bounded interval, providing a semantic way to "yield recorded.UniqueBounded
// numbers."  A seeded unique set takes in a set of seeded values, "weeds-out" any duplicate
// entries, and then yields randomly from the recorded.UniqueSeeded data.  Both types can also be given
// values after-the-fact, and this type will diligently "weed-out" duplicates.
//
//	 tl;dr - repetitive stagnation is the "number of the beast" which this type effortlessly tames
//
//		2 - recorded.Context[T any, TContext std.Instantaneous[T]]
//	 Abstract: "Give me some context around the recorded data!"
//
//	 Details:
//
// This type is quite simple but powerful: you can provide it an anonymous way to grab
// relevant context with which to wrap the yielded result.  You might think this type is
// redundant to, say, std.Data[MyContext[T]] (where MyContext[T] is your own contextual type)
// - but this type specifically REQUIRES you to provide a type which implements the
// 'std.Instantaneous[T]' interface.  JanOS is a temporal system, and the minimum 'context'
// of a temporal system should ALWAYS include a timestamp with its data points.
//
// Upon creation, you may specify if this type uses std.Data or recorded.Unique 'under the
// hood', allowing a mechanic for wrapping abstract data on Yield wile ensuring uniqueness in
// the process if desired.
//
//	 tl;dr - all JanOS data points should be paired with the contextual moments they happened
//
//		3 - recorded.Experience[T any, TContext std.Instantaneous[T]]
//	 Abstract: "Give me something new every once in a while!"
//
//	 Details:
//
// Experience is a type that LOSES its data periodically.  Let me explain - as you traverse
// through time and space, the universe will want to give you experiences.  These experiences
// may or may not leave a profound impression upon you, but those impressions typically take
// TIME to manifest.  Experience essentially provides a way to temporally 'hold' a set of
// values known to be relevant whilst providing a way to determine when that set should be
// re-seeded with new experiential values.
//
// By default, this type simply works off of a 'refractory period' of access - as long as you
// keep calling Yield the set will not clear - but once enough time has passed, the next Yield
// will clear the recorded.Experience, allowing it to record new data.  Alternatively, you may
// provide your own anonymous mechanic for determining when the set should logically clear its
// contents.
//
//	tl;dr - JanOS is designed to seed rich and diverse new experiences across time
//
// With these four types, you can begin to create human-legible semantic phrases which evaluate computationally.
// Why THESE four types?  Because all choices in life extend from the abstract concept of 'selecting' from
// existing sets - even if that set is infinity.  All data in the universe is just a number in the grand index of
// 'infinity', after all - meaning the concept of "selecting randomly unique contextual experiences" is truly all
// you need to spark off a logical train of thought in an intelligently designed system =)
//
// Words can be quite powerful, when wielded responsibly!
//
// LASTLY!  Every single one of the above types implements the recorded.Any interface, allowing these types to
// be interchangeably passed through generics.
package recorded
