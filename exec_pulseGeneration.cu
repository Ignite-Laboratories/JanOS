#include "janos.cuh"
#include <iostream>
#include <thread>
#include <chrono>
#include <cuda_runtime.h>

__global__ void countIndex(ControlStructure *control)
{
    int globalIndex = blockIdx.x * blockDim.x + threadIdx.x;

    int target = control->Thumbprint[globalIndex];
    int value = 0;
    while (control->KeepAlive)
    {
        if (value >= target)
        {
            // Toggle the signaling bit when a match is made
            control->Bits[globalIndex] ^= 1;

            // Reset the count
            value = 0;
        }
        value++;
    }
}

__global__ void walkThumbprint(ControlStructure *control)
{
    printf("Thumbprint: ");
    for (int i = 0; i < 32; i++)
    {
        printf("%d,", control->Thumbprint[i]);
    }
    printf("\n");

    while (control->KeepAlive)
    {
        countIndex<<<1, 32>>>(control);
    }
}

int main() {
    ControlStructure *d_cs = LoadControlStructure();
    walkThumbprint<<<1, 1>>>(d_cs);
    cudaDeviceSynchronize();
    return 0;
}