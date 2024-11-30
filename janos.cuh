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
    int PrimaryBit;
    int Pace;
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
    ControlStructure h_cs = {
        .KeepAlive = true,
        .Pace = 100000000
    };

    HANDLE_CUDA_ERROR(cudaMalloc((void**)&d_cs, sizeof(ControlStructure)));
    HANDLE_CUDA_ERROR(cudaMemcpy(d_cs, &h_cs, sizeof(ControlStructure), cudaMemcpyHostToDevice));
    HANDLE_CUDA_ERROR(cudaIpcGetMemHandle(&handle, d_cs));
    return handle;
}

inline ControlStructure* LoadControlStructure()
{
    cudaIpcMemHandle_t handle;
    FILE *fptr;

    if ((fptr = fopen(ipcFilePath, "r")) == nullptr) {
        printf("Error");
        return nullptr;
    }

    fread(&handle, sizeof(handle), 1, fptr);
    fclose(fptr);

    ControlStructure *device_cs = nullptr;
    HANDLE_CUDA_ERROR(cudaIpcOpenMemHandle((void**)&device_cs, handle, cudaIpcMemLazyEnablePeerAccess));
    return device_cs;
}

/**
 * KERNELS
 */

__global__ void _setPace(ControlStructure *control, const int value)
{
    control->Pace = value;
}
inline void SetPace(const int value)
{
    _setPace<<<1,1>>>(LoadControlStructure(),value);
    cudaDeviceSynchronize();
}

__global__ void _togglePrimaryBit(ControlStructure *control)
{
    int x = 0;
    do
    {
        // KEY NOTE:
        // This is how we observe faster than we increment.

        // The 'Pace' can be adjusted in real time to control
        // the master frequency of execution.
        // Key Values:
        // 0 - No throttle
        // 100000000 - ~1.3hz
        for (int z = 0; z < control->Pace; z++)
        {
            // The inner loop ALWAYS counts to INT_MAX - this is to
            // ensure observers can always observe faster than the
            // master clock can increment.
            for (int i = 0; i < INT_MAX; i++)
            {
                control->PrimaryBit ^= 1;
                printf("Toggling: %d\n", x);
                x++;
            }
        }
    } while(control->KeepAlive);
}
inline void TogglePrimaryBit()
{
    ControlStructure *control = LoadControlStructure();
    _togglePrimaryBit<<<1,1>>>(control);
    cudaDeviceSynchronize();
}

#endif