#include "janos.cuh"


[[noreturn]] int main()
{
    printf("Reading one measure\n");
    volatile ControlStructure *control = LoadControlStructure();
    int h_buffer[measureWidth];
    int* d_buffer;
    cudaMalloc(&d_buffer, measureWidth * sizeof(int));

    readOneMeasure<<<1,1>>>(control, d_buffer);

    cudaMemcpy(h_buffer, d_buffer, measureWidth * sizeof(int), cudaMemcpyDeviceToHost);

    printf("[");
    for (int i = 0; i < measureWidth; i++)
    {
        printf("%d,", h_buffer[i]);
    }
    printf("]\n");
}