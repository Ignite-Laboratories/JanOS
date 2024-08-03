# Welcome

<picture>
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/JanOS/main/enigmaneering/Logo%20-%20JanOS%20-%20Light.png">
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/JanOS/main/enigmaneering/Logo%20-%20JanOS%20-%20Dark.png">
    <img alt="JanOS Logo" src="https://raw.githubusercontent.com/Ignite-Laboratories/JanOS/main/enigmaneering/Logo%20-%20JanOS%20-%20Light.png" width="400" >
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

---

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
upon the `Engo Engine`

---

## Puzzle Piece 1 - The Oscillator

As all components operate off of a carrier signal, we must first implement a way to create such a signal.  Luckily, that
has a formulaic solution:

<picture>
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/JanOS/main/enigmaneering/Diagram%20-%20Oscillator%20-%20Circular%20Motion%20to%20Linear%20Readout_Light.png">
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/JanOS/main/enigmaneering/Diagram%20-%20Oscillator%20-%20Circular%20Motion%20to%20Linear%20Readout_Dark.png">
    <img alt="JanOS Logo" src="Diagram%20-%20Oscillator%20-%20Circular%20Motion%20to%20Linear%20Readout_Light.png" width="1000" >
</picture>

Each entity will process encoded data applied to waveforms - either by placing or perceiving anomalies in relation to
the carrier waveform. Pattern matching different ways this data is represented across time allows us to learn _when_
to perform actions or behaviors.  The entities cursor over chunks of an incoming waveform on a duty cycle of varying
width, depending on the complexity of the encoded data.

<picture>
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Ignite-Laboratories/JanOS/main/enigmaneering/Diagram%20-%20Oscillator%20-%20Anomaly_Light.png">
    <source media="(prefers-color-scheme: dark)"  srcset="https://raw.githubusercontent.com/Ignite-Laboratories/JanOS/main/enigmaneering/Diagram%20-%20Oscillator%20-%20Anomaly_Dark.png">
    <img alt="JanOS Logo" src="https://raw.githubusercontent.com/Ignite-Laboratories/JanOS/main/enigmaneering/Diagram%20-%20Oscillator%20-%20Anomaly_Light.png" width="1000" >
</picture>