#ifndef JANOS_CUH
#define JANOS_CUH

typedef uint64_t Identifier;
typedef uint64_t Instant;

extern __device__ Identifier MasterCount;
extern __device__ int Clock;

struct Dimension {
    Identifier ID;
    void *Value[];
};

struct Observation {
    Identifier ID;
    Instant HeadTime;
    void *Dimension[];
};

/**
 * @brief Initializes the JanOS environment.
 */
int Boot();

#endif