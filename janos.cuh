#ifndef JANOS_CUH
#define JANOS_CUH
#include <cuda_runtime.h>
#include <iostream>

/**
* GLOBALS
*/

inline auto ipcFilePath = "/home/rpetz/source/ignite/JanOS/ipc_handle.bin";

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
    int MeasureWidth = 64;
    int Beat = 0;
    int CoreBit;
    float Pace;
    int Bits[];
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

inline void PrintIPCHandle(cudaIpcMemHandle_t handle) {
    printf("IPC Handle: ");
    for (size_t i = 0; i < sizeof(handle); ++i) {
        printf("%02x", ((unsigned char*)&handle)[i]);
    }
    printf("\n");
}

inline void WriteToFile(const char *filename, const cudaIpcMemHandle_t& handle) {
    if (FILE *fptr = fopen(filename, "w"); fptr != nullptr)
    {
        fwrite(&handle, sizeof(handle), 1, fptr);
        fclose(fptr);
    }
}

/**
* GENERAL FUNCTIONS
*/

cudaIpcMemHandle_t CreateControlStructure()
{
    cudaIpcMemHandle_t handle;
    ControlStructure *d_cs;
    constexpr ControlStructure h_cs = {
        .KeepAlive = true,
        .Pace = 1.0
    };

    HANDLE_CUDA_ERROR(cudaMalloc(&d_cs, sizeof(ControlStructure)));
    HANDLE_CUDA_ERROR(cudaMemcpy(d_cs, &h_cs, sizeof(ControlStructure), cudaMemcpyHostToDevice));
    HANDLE_CUDA_ERROR(cudaIpcGetMemHandle(&handle, d_cs));
    return handle;
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

__global__ void _setPace(volatile ControlStructure *control, const float value)
{
    control->Pace = value;
}
inline void SetPace(const float value)
{
    _setPace<<<1,1>>>(LoadControlStructure(),value);
    cudaDeviceSynchronize();
}

__global__ void _observe(volatile ControlStructure *control)
{
    int lastState = 0;
    do
    {
        if (control->CoreBit != lastState)
        {
            lastState = control->CoreBit;
            printf("Pace: %f, Master Count: %lld, Measure Beat: %d/%d\n", control->Pace, control->MasterCount, control->Beat, control->MeasureWidth);
        }
    } while (control->KeepAlive);
}
inline void Observe()
{
    _observe<<<1,1>>>(LoadControlStructure());
    cudaDeviceSynchronize();
}

__global__ void _togglePrimaryBit(volatile ControlStructure *control)
{
    do
    {
        // KEY NOTE:
        // This is how we observe faster than we increment.

        // The 'Pace' can be adjusted in real time to control
        // the master frequency of execution.  A Pace of 1.0
        // is roughly 1hz, whereas a Pace of 0 is "as fast as
        // possible."

        int z = 0;

        for (int i = 0; i < INT_MAX; i++)
        {
            // While z is less than our pacing value...
            // Key Values:
            // 0 - No throttle
            // 10000000 - ~1hz
            if (z < 10000000 * control->Pace)
            {
                // ...keep incrementing
                z++;
            }
            else
            {
                // ...otherwise, we are ready to "step"
                z = 0;

                if (control->Beat >= control->MeasureWidth)
                {
                    control->Beat = 0;
                }
                else
                {
                    control->Beat++;
                }

                control->CoreBit ^= 1;
                control->MasterCount++;

            }
        }
    } while(control->KeepAlive);
}
inline void TogglePrimaryBit()
{
    printf("Toggling\n");
    _togglePrimaryBit<<<1,1>>>(LoadControlStructure());
    cudaDeviceSynchronize();
    printf("Terminated\n");
}

#endif