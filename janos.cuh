#ifndef JANOS_CUH
#define JANOS_CUH

typedef uint64_t Identifier;
typedef uint64_t Instant;

struct Dimension {
    Identifier ID;
    void *Value[];
};

struct Observation {
    Identifier ID;
    Instant HeadTime;
    void *Dimension[];
};

struct ControlStructure {
    int Bits[];
};

#endif