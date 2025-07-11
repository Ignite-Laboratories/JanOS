<picture>
    <source media="(prefers-color-scheme: light)" srcset="https://ignite-laboratories.github.io/assets/Logo%20-%20JanOS%20-%20Light.png">
    <source media="(prefers-color-scheme: dark)" srcset="https://ignite-laboratories.github.io/assets/Logo%20-%20JanOS%20-%20Dark.png">
    <img alt="JanOS Logo" src="https://ignite-laboratories.github.io/assets/Logo%20-%20JanOS%20-%20Light.png" width="400" >
</picture>

Welcome to JanOS - a chrono-neurological ecosystem!

Currently, the concept is centered around the idea of creating intelligent algorithms I call ARWENs.  As the project
evolves, they do as well.  This all initially started while exploring a form of binary synthesis which required a shift 
of thinking about how we currently execute computation.  This then morphed into a neurological impulse engine, and now 
I'm exploring how to build orchestrations with such a computational model.

## What's ARWEN?

Abstractly, an ARWEN is any Autonomous Robot With Ethical Navigation.  An algorithm is an ARWEN, a large language model
is an ARWEN, and when the two are combined with rich control surfaces like a nervous system and a heartbeat, you get
_**life itself**_.

### What's JanOS All About?

This project aims at a form of compression which, in reverse, can be used to synthesize anything in the known
universe.  As with DNA, an algorithm with as little as 64 bits of starting information could intermittently occupy
an infinite number of observably useful binary states.  Knowing _when_ to stop the algorithm is much less simple!

Arwen's current fugue is to use such an algorithm in reverse to find the starting conditions of a target file.

To do so one must perceive binary data as if a representation of _presence_ across _time_.  Binary is already processed
in logical widths of 8-bits called `bytes`, but they _technically_ could be read at several different widths.  For
instance, if the value of a byte is less than 64 it only needs 6-bits - less than 32, only 5.

The trouble is how do we _dynamically_ change the index bit width?

By perceiving the binary data using concepts from _**music theory**_.  If data can be read byte for byte, it's read in 
"common time" (4⁄4). If it's being read in variable widths, a new tempo must first be observed. Verses flow through 
key changes to compose movements of information.  Even then, some bytes just don't sound right when sped up and might
need extra time to resolve their note.

But _how_ do we create a system that can react to these changes instinctively?  One that can literally groove with
the language of machines?  

Well, by reflecting inward to our own neurological machinery =)

## Project Structure

This repository, itself, is a placeholder to instantiate your Go workspace.  The actual code is contained in multiple
other repositories which you should clone into the directory you cloned this one into.  As a convenience, I've included
`synchronize` scripts for CMD and Bash to both instantiate the repositories and ensure they are up to date.  Simply
execute the scripts and you're good to go!

----

**Core** - Core holds the most critical components used by all JanOS libraries - critically, this is where the impulse 
engine lives. 

- [🗎 The Neural Impulse Engine](https://github.com/Ignite-Laboratories/Enigmaneering/tree/main/enigma0)
- [🗎 Temporal Analysis](https://github.com/Ignite-Laboratories/Enigmaneering/tree/main/enigma1)

----

**Glitter** - Glitter is the artist behind all graphical rendering.

----

**Hydra** - Hydra provides a toolkit for working with concurrently parallel systems, currently acting as the window 
and graphics manager for JanOS.

----

**Support** - The "junk drawer" of JanOS:  If it's too general for the other repositories, it belongs here - everything 
from thread-safe helper functions to random generation of "Lorem Ipsum" data.

----

**Tiny** - Tiny is what facilitates binary synthesis =)

- [🗎 Binary Synthesis](https://github.com/Ignite-Laboratories/Enigmaneering/tree/main/enigma2)

----


In addition to the above repositories, there is also a non-JanOS repository documenting its creation:
### 🔗 [Enigmaneering](https://github.com/ignite-Laboratories/enigmaneering) 
The ultimate goal of this entire project is to produce a book which others can use to learn the process of exploring
abstract concepts like this.  As such, I intend to keep the documentation _thorough!_
