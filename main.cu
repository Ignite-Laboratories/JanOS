#include "janos.cuh"
#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <cuda_runtime.h>

__global__ void observe(ControlStructure *control) {
    printf("Control Bit: %d\n", control->Bits[0]);
}

int main() {
    ControlStructure h_cs = {
        .Bits = {1}
    };

    ControlStructure *d_cs;
    cudaIpcMemHandle_t handle;

    cudaError_t err = cudaMalloc((void**)&d_cs, sizeof(ControlStructure));
    if (err != cudaSuccess) {
        std::cerr << "cudaMalloc error: " << cudaGetErrorString(err) << std::endl;
        return 1;
    }

    /**err = cudaMemcpy(d_cs, &h_cs, sizeof(ControlStructure), cudaMemcpyHostToDevice);
    if (err != cudaSuccess) {
        std::cerr << "cudaMemcpy error: " << cudaGetErrorString(err) << std::endl;
        return 1;
    }*/

    err = cudaIpcGetMemHandle(&handle, d_cs);
    if (err != cudaSuccess) {
        std::cerr << "cudaIpcGetMemHandle error: " << cudaGetErrorString(err) << std::endl;
        return 1;
    }

    FILE *fptr;
    fptr = fopen("c:\\source\\ignite\\Janos\\ipc_handle.bin", "w");
    if (fptr == NULL) {
        printf("Error");
        return 1;
    }
    fwrite(&handle, sizeof(handle), 1, fptr);
    fclose(fptr);

    ControlStructure *control;
    err = cudaIpcOpenMemHandle((void**)&control, handle, cudaIpcMemLazyEnablePeerAccess);
    if (err != cudaSuccess) {
        std::cerr << "cudaIpcOpenMemHandle error: " << cudaGetErrorString(err) << std::endl;
        //return 1;
    }
    observe<<<1,1>>>(control);
    std::cout << "Press Enter to exit..." << std::endl;
    std::cin.get();

    free(&h_cs);
    cudaFree(d_cs);
    return 0;
}