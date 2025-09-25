# `E5S0 - String Wars`
### `Alex Petz, Ignite Laboratories, August 2025`

---

### Episode V - The Algorithm Strikes Back

As is traditional in all spatial operas, we begin somewhere in the middle!  This O.S. has evolved through many 
different phases, each of which needs heavy revision before _its_ introduction into the JanOS code base - _this one,_ however, 
must be written alongside the code it describes!  I made some bold claims at the start of this enigma, so let's get
started with _how_ an algorithm could even begin to make _ethically sound decisions:_  it first gets to learn how to
_describe_ the concept of "ethical soundness."  This entire enigma is devoted towards the concept of abstractly qualifying
_anything_ in a queryable fashion for intelligent decision making - again, _ethics_ and _morals._  One cannot navigate
ethically and morally unless it's psyche is foundationally built upon the concept of _defining_ whether one's actions 
are "ethical" or "moral" in nature.  This is what separates chaos from order, and what facilitated _**creation**_ through 
the natural process of _**evolution**_ =)

Before we proceed, I'd like to humbly give gratitude to two brilliant individuals, without whom this work could never have come
into creation: George Armitage Miller, who gave us the magical number 'seven' and a model for intuiting through language called 
_WordNet,_ and Christiane Fellbaum, who has continually led teams to bring that semantic lattice into existence.  We, as a
species, owe a collective debt of gratitude to the level of dedication it takes to build such a rich and powerful dataset.  _These_
incredible enigmaneers willingly put these tools into the hands of others through an insatiable drive which I greatly respect. 

_Thank you._

### Semantic Improvisation

Currently, JanOS is focused on the concept of _identity_ and _naming_ - as every single entity in the stack comes with both a
unique identifier and a string "name."  This was borne out of a basic desire: I don't want to trace 'F5D0A500' through a set
of log files when building a neural stack!  To solve this, my initial work took the brilliant NameDB from Kyle MacLeod and
paired it with the wonderful Internet SurnameDB, allowing me to randomly create _human_ names based entirely off of a real
and practical dataset.  The trouble with this is that the names aren't always, well, _simple_ - a challenge
many creatives before me typically solved through clever pairings of words and numbers.  For instance, `343 Guilty Spark` `Johnny 5`
`Hal 9000` `ED-209` or `Unit-X` - these _also_ gave a thematic element to the lineage of an identifier, with Halo's idea of
"moral adjective + abstract noun" resonating well: `Enduring Bias` and `Exuberant Witness` _immediately_ strike the tone of
a computed intelligence residing within a bounded space.

Thus, I set out on a course for producing an engine that could _generate_ a similar name on the fly: _moral adjective_ + 
_abstract noun._  Have you tried to Google how to do that?  How do you even _ask_ a computer to semantically traverse
through related lexicographical terms!?  Moreover - Halo's naming convention is occasionally, well, _dark:_  `001 Shamed Instrument` 
`2401 Penitent Tangent` and `343 Guilty Spark` to name a few!  How do I filter my adjectives and nouns to more positive 
terms, like `031 Exuberant Witness` or `686 Ebullient Prism`?  Hell, how do you ensure it could whimsically resolve that 
a fitting name should riff off the numeric value, like the author's clever choice of `000 Tragic Solitude`?

How do you even _qualify_ a _word!?_  
What makes a guitar's tone _chunky?_  
How do you correlate _mellifluousness_ against a _**data set?**_

    tl;dr - can a computer be taught what it means to stop and smell the roses?

The answer is an astoundingly resolute _yes!_  Not because it couldn't before, but because we've already given it the
ability to _temporally_ identify context - in other words, JanOS has conscious awareness of however much of its surroundings
we _give_ it access to.  This comes in the form of the vector system's temporal frame buffer, which can be considered a form 
of short-to-long term memory tuned as dense or lightweight as necessary.  This creates a kind of recursive
"Markov Chain" once you consider a vector cursor can be fed its output as an input - so long as your structures are well-formed.

So what do we _put_ on the Markov Chain?

Well - what do most instruments need, first?

_Strings!_

### Cursoring Through Prose

In JanOS, all vector types are _bounded._  For example, when creating floating point vectors they're automatically 'clamped' to
the closed interval `[0.0, 1.0]` - with clamping meaning you cannot overflow or underflow the boundaries.  A _cursor_ is a
vector that uses this mechanic to bound itself within a finite dataset.  A `Cursor[uint](0, 42)`
implies a bounded numeric value in a 42-element uint dataset which will error when you've reached the end of the slice's indices.  We don't 
need to delve too far into the weeds beyond that - instead, let's consider a _finite recursive dataset._  How does one traverse
such a structure without infinitely looping?  

Through _context._

Luckily, the `WordNet` database is a finite data set which naturally limits the branches you can follow, making it a perfect
starting point for cursoring through such a structure.  The _general concept_ is simple: lemma can either be more generalized
(a _hypernym_) or more specialized (a _hyponym_).  Lemma, in WordNet context, is the basic logical unit it works with - it can
be a word or a phrase, but the string of characters is a cohesive _unit_ -

      hyper   ←  input  →   hypo
     animal       dog       corgi
    vehicle      truck      dump truck
        ⬑  lemma  ⬏

Here, we've identified two different sets of related lemma through their hypernym and hyponym tree.  This is a _start_ towards
our goal of a "moral adjective" finder, because it allows us to traverse up and down 'generalness' to find broader sets of 
lemma.  WordNet affords us another traversal direction - _synonyms_ and _antonyms._  Starting off in these four directions, we
can theoretically take a set of starting terms and _intelligently_ traverse the entire data set to find qualitatively related lemma!  That fits
every checkmark of what I aim to achieve - so, let's start parsing the WordNet files into a logical structure that's useful to
_JanOS._