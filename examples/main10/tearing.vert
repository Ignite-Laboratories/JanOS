#version 310 es
precision mediump float;

// Input: Vertex positions in clip space
layout(location = 0) in vec2 aPos;

// Output: Pass-through texture coordinates
out vec2 fragCoord;

void main() {
    // Pass normalized device coordinates to the fragment shader
    fragCoord = aPos * 0.5 + 0.5; // Convert from [-1, 1] to [0, 1]
    gl_Position = vec4(aPos, 0.0, 1.0); // Output position in clip space
}