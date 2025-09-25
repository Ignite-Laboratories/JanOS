# Arwen's Calypso

[INSERT CARICATURE OF ARWEN]

### What's ARWEN?
Abstractly, an ARWEN is any Autonomous Robot With Ethical Navigation.  An algorithm is an ARWEN, a large language model 
is an ARWEN, and when the two are combined with rich control surfaces like a nervous system and a heartbeat, you get 
_**life itself**_.

### What's Her Calypso All About?

This project aims at a form of compression which, in reverse, can be used to synthesize anything in the known
universe.  As with DNA, an algorithm with as little as 64 bits of starting information could intermittently occupy 
an infinite number of observably useful binary states.  Knowing _when_ to stop the algorithm is much less simple!

The goal of Arwen's Calypso is to use such an algorithm in reverse to find the starting conditions of a target file.

To do so one must perceive binary data as if a representation of _presence_ across _time_.  Binary is already processed 
in logical widths of 8-bits called `bytes`, but they _technically_ could be read at several different widths.  For 
instance, if the value of a byte is less than 64 it only needs 6-bits - less than 32, only 5.  

The trouble is how do we _dynamically_ change the address width of each byte?

By perceiving the binary data using musical _**time**_.  If data can be read byte for byte, it's read in common time 
(4⁄4). If it is being read in variable widths, a new `Tempo` must first be observed. `Verses` flow through `Key Changes` 
to compose `Movements` of information.  Even then, some bytes just don't sound right when sped up and must deviate from 
the current `Time Signature`.

Luckily, musicians are pretty good at finding the _pocket!_ =)

### The Algorithm

First, let's officially define our basic structures:

`Measures` - A measure starts as 8-bits but can shrink down when it's value can be held in less bits.

`Verses` - A verse is a block of measures prepended with the details of what helped make them smaller.  A verse
also contains `Tempo` and `Time Signature` information.  If a verse could not yield any compression, its first bit
is toggled to zero and the remaining information is entirely read at full width with no prepended details.

`Movements` - A movement is a block of verses prepended with the details of what helped make them smaller.

`Timeline` - A timeline starts with some broad actions to apply to all bytes before grouping them into movements.

`Keys` and `Key Changes` - Both are 3-bit patterns that are XORd against the 3 most significant bits of a measure. 
The goal is to give multiple opportunities to toggle _these_ 1s to a 0.  When you do so effectively, the byte can then 
be represented in fewer bits.  The act of performing a key change is known as `Transposing` the measure.

`Inversion` - Inversion is the act of XORing all 1s against binary data, which toggles each position. Binary, naturally, 
exists in one of two states: predominantly light, or predominantly dark.  An inversion operation attempts to 
make the upcoming measures as _light_ as possible, thus, holding their smallest value.

`Tempo` - The tempo states whether the current measures are being reduced by 2 bits (0) or 3 bits (1).

`Time Signatures` - A time signature is used to dictate which measures should or should not follow the current tempo.
It is represented as a bitmap of the upcoming measures, with a 0 indicating _not_ to speed the measure up.