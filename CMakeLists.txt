cmake_minimum_required(VERSION 3.29)
set(CMAKE_CUDA_ARCHITECTURES "native")
set(CMAKE_CUDA_COMPILER /usr/local/cuda/bin/nvcc)
set(CMAKE_CUDA_STANDARD 20)

# CREATE CONTROL STRUCTURE
project(CreateControlStructure CUDA)
add_executable(CreateControlStructure exec_createControlStructure.cu)
set_target_properties(CreateControlStructure PROPERTIES CUDA_SEPARABLE_COMPILATION ON)

# TOGGLE PRIMARY BIT
project(TogglePrimaryBit CUDA)
add_executable(TogglePrimaryBit exec_togglePrimaryBit.cu)
set_target_properties(TogglePrimaryBit PROPERTIES CUDA_SEPARABLE_COMPILATION ON)

# OBSERVER
project(Observer CUDA)
add_executable(Observer exec_observe.cu)
set_target_properties(Observer PROPERTIES CUDA_SEPARABLE_COMPILATION ON)
