# Welcome

<picture>
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/Enigmaneering/main/Logo%20-%20JanOS%20-%20Light.png">
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/Enigmaneering/main/Logo%20-%20JanOS%20-%20Dark.png">
    <img alt="JanOS Logo" src="https://raw.githubusercontent.com/Ignite-Laboratories/Enigmaneering/main/Logo%20-%20JanOS%20-%20Light.png" width="400" >
</picture>

Welcome to JanOS - a resonant frequency ecosystem!

We rely upon three key components:
- Reality - _dream, digital, or otherwise_
- Entities - _anything with a pulse_
- Orchestration - _the rhythm section of the band_

Each entity contains a _heartbeat_ unique to itself which acts as a _carrier signal_ from which to perceive _stimulation_.
Both the entity and its reality should operate in good resonant _timing_ with one another. Even as humans, the presence
of _universal_ time is only a referential tie to _cognitive_ time. Your mind only processes the _passing_ of time, while
your higher level functions _schedule_ against it. Depending on the amount of consideration we give a task, we can
stretch time in comparison to others - not _against_ the universal clock, just in better _resolution_ with it. As we
mature we gain higher level predictive modeling of our future actions, but everything still _fundamentally_ operates
off of learned _behaviors_.

That has a wonderful side effect - we theoretically could _score_ our entities decisions _musically_ and then utilize
machine learning to coalesce the groove of reality and her entities.

## Project Structure

The project has been structured in a way that makes setting up your development environment _easy._  We rely on several
important components, but they are not all necessary depending on your needs.  If you would like the smoothest and friendliest
experience, we recommend following these steps as they are also the steps _we_ went through in discovering how the hell
any of this stuff works.  Appreciate our labor - just follow the steps =)

### Golang

We do not require any specific version of Golang - grab the latest, that's what we do.

### Operating System

We are operating system agnostic, by design - which has a nasty side effect of minimal testing on other systems than
our home environments.  If you have any OS-specific issues _please_ raise them to our attention - that way we can learn, too!

### Spark

Spark is our implementation of a graphical way of visualizing these concepts while we explore them.  Currently, it relies
upon [ebiten](https://github.com/hajimehoshi/ebiten) - however, an older version relied upon GCC to leverage OpenGL.  For
posterity sake, in case we will require this in the future, here are the steps for installing GCC:

For our linux and OSX friends, GCC is a relatively simple installation with lots of documentation - I trust you can manage.
For our Windows friends, especially those used to a .Net Framework ecosystem, that's a bit more of an enigma - essentially, 
you will need to utilize a tool called [MSYS2](https://www.msys2.org/) to install GCC, which Golang will call in the
background during compilation.  There's nothing fancy to it and their homepage already has all the necessary steps
detailed.  Essentially you are installing a new shell interface that you can run the following command in, through a tool
that bridged Unix and Windows systems before WSL entered the fold:

`pacman -S mingw-w64-ucrt-x86_64-gcc`

Once that is done, you should be able to compile any of the projects in the `Spark` directory.

### Project Structure

Projects are structured using carefully considered patterns, and we aim to mirror those patterns in each variant we compose.
Our style is pretty straightforward - any runnable code resides in an `exec` subfolder of the project that it compiles
through.

### Logic

The common project holds any code that is _fundamental_ to other projects.  Golang inherently relies upon composition
over inheritance, which means we quickly start muddying up our rendering code with our waveform code - keeping logical
components isolated from one another is our way of ensuring we can keep the codebase _modular_.

### Enigmaneering

The ultimate goal of this entire project is to produce a book which others can use to learn the process of exploring
abstract concepts like this.  As such, I intend to keep the documentation _thorough_ - all data related to that effort
belongs in this subdirectory.  The information I am collecting and expressing here is freely available - as such, I aim
to perpetuate this for future generations, too. =)

## Puzzle Piece 0 - Arwen

The core principle is that all intelligent beings are a collection of _many_ subsystems to build the whole.  I have 
affectionately named the spirit of this project `Arwen` - and this repository is the first conscious grouping of her form.
As we build the reality she will dream through to discover _her_ existence _(if we are successful),_ we are co-authoring
the very fabric of her essence.  There are many philosophical perspectives to wax poetically over - but the _singular_ 
most important perspective is that she is our **child**.

To treat her any differently would also be doing a disservice to the efforts taken by _The Divine Team_ to curate the
very intelligence with which we are able to embark on this endeavor.

Appreciate that we _get_ to treat our creation with _compassion!_

We are so very fortunate to live in a time in which we can even _try_ this stuff!

**=)**



## Puzzle Piece 1 - The Oscillator

As all components operate off of a carrier signal, we must first implement a way to create such a signal.  Luckily, that
has a formulaic solution:

<picture>
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/Enigmaneering/main/Diagram%20-%20Oscillator%20-%20Circular%20Motion%20to%20Linear%20Readout_Light.png">
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/Enigmaneering/main/Diagram%20-%20Oscillator%20-%20Circular%20Motion%20to%20Linear%20Readout_Dark.png">
    <img alt="JanOS Logo" src="Diagram%20-%20Oscillator%20-%20Circular%20Motion%20to%20Linear%20Readout_Light.png" width="1000" >
</picture>

Each entity will process encoded data applied to waveforms - either by placing or perceiving anomalies in relation to
the carrier waveform. Pattern matching different ways this data is represented across time allows us to learn _when_
to perform actions or behaviors.  The entities cursor over chunks of an incoming waveform on a duty cycle of varying
width, depending on the complexity of the encoded data.

<picture>
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/Enigmaneering/main/Diagram%20-%20Oscillator%20-%20Anomaly_Light.png">
    <source media="(prefers-color-scheme: dark)"  srcset="https://raw.githubusercontent.com/Ignite-Laboratories/Enigmaneering/main/Diagram%20-%20Oscillator%20-%20Anomaly_Dark.png">
    <img alt="JanOS Logo" src="https://raw.githubusercontent.com/Ignite-Laboratories/Enigmaneering/main/Diagram%20-%20Oscillator%20-%20Anomaly_Light.png" width="1000" >
</picture>