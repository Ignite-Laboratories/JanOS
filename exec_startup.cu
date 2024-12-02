#include "janos.cuh"

void printIPCHandle(cudaIpcMemHandle_t handle) {
    printf("IPC Handle: ");
    for (size_t i = 0; i < sizeof(handle); ++i) {
        printf("%02x", ((unsigned char*)&handle)[i]);
    }
    printf("\n");
}

void writeToFile(const char *filename, const cudaIpcMemHandle_t& handle) {
    if (FILE *fptr = fopen(filename, "w"); fptr != nullptr)
    {
        fwrite(&handle, sizeof(handle), 1, fptr);
        fclose(fptr);
    }
}

[[noreturn]] int main() {
    printf("JanOS\n");

    cudaIpcMemHandle_t handle;
    ControlStructure *d_cs;
    constexpr ControlStructure h_cs = {
        .KeepAlive = true,
    };

    HANDLE_CUDA_ERROR(cudaMalloc(&d_cs, sizeof(ControlStructure)));
    HANDLE_CUDA_ERROR(cudaMemcpy(d_cs, &h_cs, sizeof(ControlStructure), cudaMemcpyHostToDevice));
    HANDLE_CUDA_ERROR(cudaIpcGetMemHandle(&handle, d_cs));

    printIPCHandle(handle);
    writeToFile(ipcFilePath, handle);

    toggle<<<1,1>>>(d_cs);
    printf("Toggling Core Bit\n");
    cudaDeviceSynchronize(); // Implicitly block until KeepAlive is false

    printf("Cleaning Up\n");
    cudaIpcCloseMemHandle(&handle);
    cudaFree(d_cs);
    printf("Goodbye\n");
}