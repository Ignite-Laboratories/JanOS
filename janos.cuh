#ifndef JANOS_CUH
#define JANOS_CUH
#include <cuda_runtime.h>
#include <iostream>

/**
* GLOBALS
*/

inline auto ipcFilePath = "/home/rpetz/source/ignite/JanOS/ipc_handle.bin";
constexpr int measureWidth = 4096;

/**
* DATA STRUCTURES
*/

typedef long long Identifier;
typedef long long Instant;

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
    bool KeepAlive;
    long long MasterCount;
    int CoreBit;
    int Beat;
    int Buffer[measureWidth];
};

/**
* SUPPORT
*/

#define HANDLE_CUDA_ERROR(err) (HandleError(err, __FILE__, __LINE__))
inline void HandleError(const cudaError_t err, const char *file, const int line) {
    if (err != cudaSuccess) {
        std::cerr << "CUDA Error: " << cudaGetErrorString(err) << " in file " << file << " at line " << line << std::endl;
        exit(EXIT_FAILURE);
    }
}

inline volatile ControlStructure* LoadControlStructure()
{
    cudaIpcMemHandle_t handle;
    FILE *fptr;

    if ((fptr = fopen(ipcFilePath, "r")) == nullptr) {
        printf("Error");
        return nullptr;
    }

    fread(&handle, sizeof(handle), 1, fptr);
    fclose(fptr);

    volatile ControlStructure *device_cs = nullptr;
    HANDLE_CUDA_ERROR(cudaIpcOpenMemHandle((void**)&device_cs, handle, cudaIpcMemLazyEnablePeerAccess));
    return device_cs;
}

/**
 * KERNELS
*/

__global__ void toggle(volatile ControlStructure *control)
{
    do
    {
        // KEY NOTE:
        // This is how we observe faster than we increment.
        int x = 0;
        for (int i = 0; i < INT_MAX; i++)
        {
            if (x == 40000000)
            {
                control->CoreBit ^= 1;
                control->MasterCount++;
                control->KeepAlive = true;
                x = 0;
            }
            x++;
        }
    } while(control->KeepAlive);
}

__global__ void loop(volatile ControlStructure *control)
{
    int lastState = control->CoreBit;
    do
    {
        if (control->CoreBit != lastState)
        {
            lastState == control->CoreBit;

            control->Beat++;
            if (control->Beat == measureWidth)
            {
                printf("Looping\n");
                control->Beat = 0;
            }
        }
    } while (control->KeepAlive);
}

__global__ void readOneMeasure(volatile ControlStructure *control, int* output)
{
    // Wait until one measure has passed...
    int end = control->Beat - 1;
    if (end < 0)
    { end = measureWidth; }
    printf("Waiting until %d\n", end);


    while (control->Beat != end)
    {
    }
    printf("Done!\n");

    // ...then copy the buffer and return control to the host
    for (int i = 0; i < measureWidth; i++)
    {
        output[i] = control->Buffer[i];
    }
}

__global__ void writeOnBeat(volatile ControlStructure *control, int beat)
{
    bool gotValue = false;
    do
    {
        if (control->Beat == beat && !gotValue)
        {
            control->Buffer[beat] = control->CoreBit;
            gotValue = true;
        }
        else
        {
            gotValue = false;
        }
    } while (control->KeepAlive);
}

__global__ void shutdown(volatile ControlStructure *control)
{
    control->KeepAlive = false;
}

#endif