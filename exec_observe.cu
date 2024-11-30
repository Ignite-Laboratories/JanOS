#include "janos.cuh"
#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <cuda_runtime.h>

__global__ void detectHigh(ControlStructure *control) {
    bool locked = false;
    while (true)
    {
        if (control->PrimaryBit == 0)
        {
            locked = false;
        }
        else if (!locked)
        {
            locked = true;
            printf("Trigger\n");
        }
    }
}

__device__ int lock = 0;
__global__ void wait(ControlStructure *control, int value)
{
    while (control->PrimaryBit != value && lock == 1)
    {
        lock = 0;
    }
    lock = 1;
}

int main() {
    ControlStructure *d_cs = LoadControlStructure();

    //detectHigh<<<1,1>>>(d_cs);
    while (true)
    {
        wait<<<1,1>>>(d_cs, 1);
        cudaDeviceSynchronize();
        printf("Triggered\n");
    }
    return 0;
}