#include "janos.cuh"


[[noreturn]] int main()
{
    printf("Launching %d observers\n", measureWidth);
    volatile ControlStructure *control = LoadControlStructure();

    for (int i = 0; i < measureWidth; i++)
    {
        writeOnBeat<<<1,1>>>(control, i);
    }

    cudaDeviceSynchronize(); // Implicitly block until KeepAlive is false
}