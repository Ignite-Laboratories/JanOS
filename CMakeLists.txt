cmake_minimum_required(VERSION 3.29)
project(JanOS CUDA)

set(CMAKE_CUDA_STANDARD 20)

add_executable(JanOS main.cu
        janos.cuh)
add_executable(Observer main2.cu
        janos.cuh)

set_target_properties(JanOS PROPERTIES
        CUDA_SEPARABLE_COMPILATION ON)
