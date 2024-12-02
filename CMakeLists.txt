cmake_minimum_required(VERSION 3.29)
set(CMAKE_CUDA_ARCHITECTURES "native")
set(CMAKE_CUDA_COMPILER /usr/local/cuda/bin/nvcc)
set(CMAKE_CUDA_STANDARD 20)

project(Startup CUDA)
add_executable(Startup exec_startup.cu)
set_target_properties(Startup PROPERTIES CUDA_SEPARABLE_COMPILATION ON)

project(LoopingIncrementer CUDA)
add_executable(LoopingIncrementer exec_loopingIncrementer.cu)
set_target_properties(LoopingIncrementer PROPERTIES CUDA_SEPARABLE_COMPILATION ON)

project(WriteOnBeat CUDA)
add_executable(WriteOnBeat exec_writeOnBeat.cu)
set_target_properties(WriteOnBeat PROPERTIES CUDA_SEPARABLE_COMPILATION ON)

project(ReadOneMeasure CUDA)
add_executable(ReadOneMeasure exec_readOneMeasure.cu)
set_target_properties(ReadOneMeasure PROPERTIES CUDA_SEPARABLE_COMPILATION ON)

project(Shutdown CUDA)
add_executable(Shutdown exec_shutdown.cu)
set_target_properties(Shutdown PROPERTIES CUDA_SEPARABLE_COMPILATION ON)
