#include "janos.cuh"
#include <chrono>

#include <cstdio>
#include <windows.h>

__device__ Identifier MasterCount;
__device__ int Clock;

__global__ void test() {
    int index = blockIdx.x * blockDim.x + threadIdx.x;
}

__global__ void trigger(int state) {
    bool triggered = false;
    clock_t start = clock();
    clock_t now;
    for (;;) {
        now = clock();
        clock_t cycles = now > start ? now - start : now + (0xffffffff - start);

        if (Clock == state && !triggered) {
            triggered = true;
            printf("%d\n", state);
        }

        if (Clock != state) {
            triggered = false;
        }

        if (cycles > 100000) {
            break;
        }
    }

}

__global__ void pulse() {
    Clock ^= 1;
}

void watch1() {
    trigger<<<1,1>>>(0);
}

void watch2() {
    trigger<<<1,1>>>(1);
}

int Boot() {
    cudaStream_t low;
    cudaStreamCreate(&low);

    cudaStream_t high;
    cudaStreamCreate(&high);

    trigger<<<1,1,0,low>>>(0);
    trigger<<<1,1,0,high>>>(1);

    cudaStream_t p;
    cudaStreamCreate(&p);

    while(true) {
        auto start = std::chrono::high_resolution_clock::now();
        pulse<<<1,1,0,p>>>();
        cudaStreamSynchronize(p);

        auto end = std::chrono::high_resolution_clock::now();
        std::chrono::duration<double, std::milli> elapsed = end - start;
        int sleep_duration = 500 - static_cast<int>(elapsed.count());
        Sleep(sleep_duration);
    }

    cudaStreamDestroy(low);
    cudaStreamDestroy(high);
    cudaStreamDestroy(p);
    return 0;
}