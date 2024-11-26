#include "janos.cuh"
#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <cuda_runtime.h>

__global__ void observe(ControlStructure *control) {
    printf("Control Bit: %d", control->Bits[0]);
}

int main() {
    cudaIpcMemHandle_t handle;
    FILE *fptr;

    if ((fptr = fopen("c:\\source\\ignite\\Janos\\ipc_handle.bin", "r")) == NULL) {
        printf("Error");
        return 1;
    }

    fread(&handle, sizeof(handle), 1, fptr);
    fclose(fptr);


    ControlStructure *control;
    cudaIpcOpenMemHandle((void**)&control, handle, cudaIpcMemLazyEnablePeerAccess);
    observe<<<1,1>>>(control);
    return 0;
}