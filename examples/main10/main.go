package main

import (
	_ "embed"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"runtime"
	"time"
)

func init() {
	// This locks the Go main thread to the OS thread required for OpenGL
	runtime.LockOSThread()
}

//go:embed tearing.frag
var FragmentShader string
var FragmentShaderID uint32

//go:embed tearing.vert
var VertexShader string
var VertexShaderID uint32

var Program uint32

func main() {
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("failed to initialize SDL: %v", err)
	}
	defer sdl.Quit()

	// Create SDL window with OpenGL
	window, err := sdl.CreateWindow("Screen Tearing Test",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1024, 768,
		sdl.WINDOW_OPENGL|sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		log.Fatalf("failed to create window: %v", err)
	}
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	// Create OpenGL context
	_, err = window.GLCreateContext()
	if err != nil {
		log.Fatalf("failed to create OpenGL context: %v", err)
	}

	if err := sdl.GLSetSwapInterval(-1); err != nil {
		log.Printf("Adaptive V-Sync not available, falling back to V-Sync: %v", err)
		if err := sdl.GLSetSwapInterval(1); err != nil {
			log.Printf("Standard V-Sync also failed: %v", err)
		}
	}

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalf("failed to initialize OpenGL: %v", err)
	}

	// Compile shaders and link program
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	csources, free := gl.Strs(VertexShader + "\x00")
	gl.ShaderSource(vertexShader, 1, csources, nil)
	free()
	gl.CompileShader(vertexShader)

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	csources, free = gl.Strs(FragmentShader + "\x00")
	gl.ShaderSource(fragmentShader, 1, csources, nil)
	free()
	gl.CompileShader(fragmentShader)

	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	gl.UseProgram(shaderProgram)

	// Main loop
	running := true

	vertices := []float32{
		-1.0, -1.0, // Bottom-left corner
		1.0, -1.0, // Bottom-right corner
		-1.0, 1.0, // Top-left corner
		-1.0, 1.0, // Top-left corner
		1.0, -1.0, // Bottom-right corner
		1.0, 1.0, // Top-right corner
	}

	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Sym {
					case sdl.K_ESCAPE:
						running = false
					}
				}
			}
		}

		gl.ClearColor(0.25, 0.25, 0.25, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		elapsed := float32(time.Since(core.Inception).Seconds())
		timeUniform := gl.GetUniformLocation(shaderProgram, gl.Str("time\x00"))
		gl.Uniform1f(timeUniform, elapsed)

		// Bind the VAO and draw the fullscreen quad
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		gl.BindVertexArray(0)

		// Swap the buffers
		window.GLSwap()

	}

	// Cleanup
	gl.DeleteVertexArrays(1, &vao)
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	gl.DeleteProgram(shaderProgram)
	window.Destroy()
	sdl.Quit()
}
