#include "janos.cuh"

[[noreturn]] int main()
{
    printf("Looping\n");
    volatile ControlStructure *control = LoadControlStructure();
    loop<<<1,1>>>(control);
    cudaDeviceSynchronize(); // Implicitly block until KeepAlive is false
    printf("Goodbye\n");
}