#version 310 es
precision mediump float;

// Uniforms
uniform float time;

// Output color
out vec4 fragColor;
in vec2 fragCoord;

void main() {
    vec2 uv = fragCoord;

    // Calculate horizontal sine wave effect
    float f = abs(uv.x - 0.5 + sin(time * 5.0) * 0.4) < 0.1 ? 1.0 : 0.0;

    // Output the color (white = 1.0, black = 0.0)
    fragColor = vec4(f, f, f, 1.0); // Grayscale (R, G, B are the same)
}
