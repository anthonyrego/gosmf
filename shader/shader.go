package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var state struct {
	shader *Shader
}

// GetActive get the current shader in the state
func GetActive() *Shader {
	return state.shader
}

// Shader is a struct for render programs
type Shader struct {
	vert       string
	frag       string
	id         uint32
	loaded     bool
	Projection int32
	Camera     int32
	Model      int32
}

// Activate will draw everything with this program and load it if it isnt loaded.
func (shader *Shader) Activate() {
	if !shader.loaded {
		program, err := newProgram(shader.vert, shader.frag)
		if err != nil {
			panic(err)
		}
		shader.id = program
		shader.loaded = true

		gl.UseProgram(shader.id)

		shader.Projection = gl.GetUniformLocation(program, gl.Str("projection\x00"))
		shader.Camera = gl.GetUniformLocation(program, gl.Str("camera\x00"))
		shader.Model = gl.GetUniformLocation(program, gl.Str("model\x00"))

	}
	state.shader = shader
	gl.UseProgram(shader.id)
}

// Use will load a default shader from a map of shaders
func Use(shaderName string) (shader *Shader) {
	if shader := defaultShaders[shaderName]; shader.vert != "" {
		shader.Activate()
		return &shader
	}
	return nil
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources := gl.Str(source)
	gl.ShaderSource(shader, 1, &csources, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
