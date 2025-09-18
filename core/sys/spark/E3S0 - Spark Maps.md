# `E3 - Any Maps`
### `Alex Petz, Ignite Laboratories, September 2025`

---

Parseltongue
---

Before we dive into the concept of a DNA  file, I get to talk about the concept of 'parsing' data.  From a very
early point in this process I thought that all data needed structure and type, but eventually I refined my code
until I recognized a simple pattern: I wanted to encode _verbs,_ not _nouns._  My goal shifted quickly
from the idea of storing _data_ to the concept of storing _actions._  In that, I began to take in a variadic of
type `any` before introspecting it to determine what the user might want to do.  That
wound up being a very powerful concept because it simplified my APIs from looking like this:

    num.ParseIntFromString(string) int
    num.ParseFloat64FromString(string) float64
    num.ParseFloat32FromString(string) float32

    num.ParseIntFromBytes([]byte) int
    num.ParseFloat64FromBytes([]byte) float64
    num.ParseFloat32FromBytes([]byte) float32

    num.ParseIntFromReader(io.Reader) int
    num.ParseFloat64FromReader(io.Reader) float64
    num.ParseFloat32FromReader(io.Reader) float32

...to this:

    num.ParseFromString[T num.Primitive](string) T
    num.ParseFromBytes[T num.Primitive]([]byte) T
    num.ParseFromReader[T num.Primitive](io.Reader) T

...and finally this:

    spark.Parse[T any](any, ...hint.Hint) T

Legitimately - THIS is the path I took!  Everything started from a frustration that I could not compare two numeric
types without casting one to another.  I lived the movie "Groundhog Day" on this concept until I eventually realized 
I could parse _anything_ into _anything_ - rather than explicitly _caring_ what I took in.  From there, I began to look around
at the inner workings of how existing systems communicated with each other.  What I realized was they all spoke the same
language, just with different _verbs._  For example - JSON "quotes" data, while XML "tags" elements. SQL "queries" data, 
while ZIP "compresses" it.  Marshallers "serialize" data, while algorithms "calculate" a result. 

All _wildly_ different, but abstractly similar - verbs which _identify_ some perspective of logic.  

    tl;dr - an intelligent parser, such as a Human, can identify many perspectives of the same input

Because of this, I was able to distill the concept of "identified" data into a very simple format, which
I call an "Any Map."  Yes, I already hear the groans - "what's wrong with JSON!?" - well, nothing!  In
fact, Spark can parse JSON (and many others).  I've taken a very different approach, however - the data
is _off-limits_ from requiring modification to fit the schema!  For instance, JSON requires you escape 
quote characters - complicating the interpreter and serializer.  Instead, any maps capture the data's 
boundaries, and the delimiter for each key can change to accomodate the data intuitively.

The idea (as with many modern formats) is that everything should be _Human-legible_, while retaining _Machine-readability_.
It should be parsable by literally any intelligent entity reading the damn data!  `READ: A Human OR a Machine.`  That led 
me to the following core tenants of the "any map" format:

- 0 It should always be in Human-readable form
- 1 A "compact" form should only shrink out Human-friendly whitespace without affecting the stored data
- 2 It should retain the EXACT Human-friendly formatting written by the author and NEVER auto-format to anyone else's standards unless compacted
- 3 There should be no restrictions over what delimiters are used, aside from newlines and comments 

In that, I immediately drafted the below specification:

The Any Map Schema
---

An "any map" is a Unicode schema which contains only one logical element: an **entry**.  Here's a few examples:

    :keyA:valueA:
    :keyB:valueB:
    :keyC:valueC:
    :keyD:valueD:

An entry consists of three **components** - the **key**, **value**, and a **delimiter** character.  The **data** of the entry pertains
to either the key or the value independently of the delimiter. The delimiter character is the first non-whitespace character 
outside a comment and can be any single Unicode character. This means the below entries are all valid:

    @keyA@valueA@
    #keyB#valueB#
    $keyC$valueC$
    %keyD%valueD%

The colloquial delimiter is `:` and the only rule is that the chosen delimiter cannot exist within the entry's data.

Whitespace within a comment or entry's data is considered untouchable, but all other whitespace is entirely 
optional and ignored by interpreters.  This means you can store things "compactly":

    @keyA@valueA@#keyB#valueB#$keyC$valueC$%keyD%valueD%

Comments are treated as such:

    // "Slash Slash" comments terminate with a newline character
    /* "Star" comments can be used inline or as multi-line blocks */

Spark maps can be nested within each other by using double-delimiter entries.  This treats the 'value' as
a new spark map:

    ::Network::
        :address:192.168.1.11:
        :port:4242:
    ::

    // In compact form:
    ::Network:::address:192.168.0.42::port:4242:::

You're welcome to mix delimiters dynamically when encountering an entry whose data requires the current delimiter:

    ::Network::
        :hostname:ignitelabs.net:
        |address|192.168.1.11:4242|
    ::

Next, you can recursively nest spark maps inside other spark maps - as long as the inner map uses a different double-delimiter:

    ::Configuration::
        %Description%Multi-line data must not
    introduce any unintended whitespace by the author%
        %Keys have the
    same rules as values%That's just how the "boundary" formatting principle works =)%

        &&Display&&
            :width:1024:
            :height:768:

            **Resolutions**
                %%0%%
                    :width:1024:
                    :height:768:
                %%
                %%1%%
                    :width:640:
                    :height:480:
                %%
            **
        &&

        &&Network&&
            :hostname:ignitelabs.net:
            |address|192.168.1.11:4242|
        &&

        **Peripherals & Devices**
            :mouse:true:
            :keyboard:true:
        **
    ::

    // In compact form:
    ::Configuration::&&Display&&:width:1024::height:768:**Resolutions**%%0%%:width:1024::height:768:%%%%1%%:width:640::height:480:%%**&&&&Network&&:hostname:ignitelabs.net:|address|192.168.1.11:4242|&&**Peripherals & Devices**:mouse:true::keyboard:true:**::

That's really it!  You may wonder why I chose not to put any method of 'escaping' the chosen delimiter?  Well, because this is meant
to represent a map - not to serialize the data.  The only logical way I can think of intuitively escaping a delimiter is through doubling 
it up, which conflicts with the signal for nesting.  I feel strongly that this limitation is a blessing, as it aims to keep spark maps 
from getting TOO deep.

Structure Parsing
---

The data is inherently typeless.  For instance, quotes are not required to store strings, only a logical delimiter that can identify
the bounds of an entry's value.  This is _especially_ important as systems designed to interop with each other already know how
to do so - this is just a _substrate_ for passing a message between them.  That means parsing spark map data into a logical structure
is as simple as the following:

    type MyStruct struct {
        Configuration string
    }

    config := spark.Parse[myStruct](":Configuration:data:")

If given a string (or a functional type that yields a string) spark will map it onto a structure automatically.  If you have a more 
complex structure that requires custom parsing, you can **optionally** implement a Parse method on the target structure like such:

    type MyStruct struct {
        Data CustomType
    }

    func (m MyStruct) Parse(key string, value any) bool {
        switch key {
        case "Data":
            ...set m.Data here
            return true
        default:
            return false
        }
    }

If present, spark will call the Parse method for every entry in the level of the map it's currently processing.  If
you return `true` spark will skip that value and consider it 'parsed' otherwise it will attempt to coerce the value itself.
If it fails to do so, spark will intentionally hard panic to alert you of the issue for immediate resolution.  Spark can handle
structure fields in the target type as well, recursively dropping into a sublevel of the spark map.

    type MySubStruct struct {
        Configuration string
    }

    type MyStruct struct {
        Data MySubStruct
    }

    data := "::Data::
                :Configuration:Value:
             ::"

    config := spark.Parse[MyStruct](data)

Hints
---

What if your source data is JSON?  Or XML?  Or, well, _anything?_  This is where 'hints' come into play, as they explicitly define what
languages spark currently speaks in.  You can see what hints are currently available from the `hint` enumeration package, and each is well-documented.

    type MyStruct struct {
        Configuration string
    }

    config := spark.Parse[myStruct]("{\"Configuration\":\"data\"}", hint.JSON)

Primitive Types
---
Spark does NOT require you to request a structure type, though!  Let's say you want to parse a numeric type from a string:

    float := spark.Parse[float64]("-123.456")

Because this requests a primitive Go type, it parses strings (or string providers) as if they hold a raw numeric value.
A primitive Go type is considered anything but composite or slice types, with the only exception being complex numbers
(which are a composite primitive type).  I call these types `primitive` since they form the foundational building blocks 
of all structures above (as every super type derives off of pointer technology).  