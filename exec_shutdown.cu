#include "janos.cuh"

[[noreturn]] int main() {
    printf("Shutting Down\n");
    volatile ControlStructure *control = LoadControlStructure();
    shutdown<<<1,1>>>(control);
    cudaDeviceSynchronize();
}