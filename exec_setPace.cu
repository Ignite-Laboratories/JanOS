#include "janos.cuh"

int main(int argc, char* argv[]) {
    float pace = std::atof(argv[1]);
    printf("Setting pace to %f", pace);
    SetPace(pace);
    return 0;
}
