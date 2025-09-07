package see

// Neuron
/*
# " 𝑁𝑒𝑢𝑟𝑜𝑛𝑠 "

A neuron is any impulsable type that can reveal a number - abstractly speaking!  A neuron must provide two methods:

	type Neuron interface {
	  Impulse()
	  Reveal() any
	}

𝐼𝑚𝑝𝑢𝑙𝑠𝑒() and 𝑅𝑒𝑣𝑒𝑎𝑙() both spark the underlying neurological action, but 𝑅𝑒𝑣𝑒𝑎𝑙() typically invokes a sequential step beyond
the impulse action.  Neurological activity typically can be broken down into a standardized "lifecycle":

	0 - Observe a potential
	1 - Describe the observation
	2 - Participate in calculation

This might sound familiar to those of us who've dabbled in Dialectic Behavioral Therapy.  This standard 𝑂𝐷𝑃 loop is what
drives, as I hope to inspire, all neurological execution - including yours!  Neurons are microscopic programmatic
loops that observe a threshold before potentially describing their observation to participate in a neural network.

	tl;dr - observe, describe, participate

At its core, this can best be described through the concept of a 'realized number'.  Fundamentally, since a number can
only be observed as 'periodic' or 'irrational' in the context of a mathematical operation, it cannot be passed around
in a mutable fashion. Thus, who's to say if a number is "irrational" other than the number, itself, or its creator?
This is the very basic concept that drives the realized number: the ability to -create- the number is "baked" right in!
When you 𝐼𝑚𝑝𝑢𝑙𝑠𝑒() a realized number, it checks a potential before describing its observed result - then, if asked to
𝑅𝑒𝑣𝑒𝑎𝑙() itself, participates by printing a result.

I hope my hyperbole has landed well - it's driven neurologically, just as YOU are in exploring this documentation!
The only difference is the richness of the questions your psyche's "potentials" have yet to 𝑅𝑒𝑣𝑒𝑎𝑙() - just as the
creation of this operating system has done for myself.

So, let's circle back to the neural interface - 𝐼𝑚𝑝𝑢𝑙𝑠𝑒(), and 𝑅𝑒𝑣𝑒𝑎𝑙().  Throughout much of this project I had spun
around the term 'yield' for the reveal method - but I wanted to instill that a number is not submitting to a higher
power!  Instead, it 'reveals' its current artistic identity =)

Genuinely - I could write for days on this subject, but it's that simple, folks!

	tl;dr - your impulses are revealed by the logical operating system YOUR psyche created to survive.

The systems which neural architectures construct are NOT complex!  Don't be afraid to dive in, the water's fine =)

see.ActionPotentials
*/
type Neuron byte
