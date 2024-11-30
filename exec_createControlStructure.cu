#include "janos.cuh"
#include <unistd.h>

[[noreturn]] int main() {
    const cudaIpcMemHandle_t handle = CreateControlStructure();

    PrintIPCHandle(handle);
    WriteToFile(ipcFilePath, handle);

    // Keep alive
    while (true)
    {
        sleep(1);
    }
}