package render

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
)

// Shader is a struct for render programs
type Shader struct {
	vert       string
	frag       string
	id         uint32
	loaded     bool
	projection int32
	camera     int32
	model      int32
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

		shader.projection = gl.GetUniformLocation(program, gl.Str("projection\x00"))
		shader.camera = gl.GetUniformLocation(program, gl.Str("camera\x00"))
		shader.model = gl.GetUniformLocation(program, gl.Str("model\x00"))

		// Move des outta har
		project := mgl32.Ortho2D(0, 800, 600, 0)
		gl.UniformMatrix4fv(shader.projection, 1, false, &project[0])
		cam := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
		gl.UniformMatrix4fv(shader.camera, 1, false, &cam[0])
	}
	state.shader = shader
	gl.UseProgram(shader.id)
}

// UseShader will load a default shader from a map of shaders
func UseShader(shaderName string) (shader *Shader) {
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
