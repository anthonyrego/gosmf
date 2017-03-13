package shader

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var shaderList = map[string]Shader{}

var state struct {
	shader *Shader
}

// GetActive get the current shader in the state
func GetActive() *Shader {
	return state.shader
}

// Shader is a struct for render programs
type Shader struct {
	vert     string
	frag     string
	id       uint32
	loaded   bool
	uniforms map[string]int32
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

		for key, _ := range shader.uniforms {
			shader.uniforms[key] = gl.GetUniformLocation(program, gl.Str(key+"\x00"))
		}
	}
	state.shader = shader
	gl.UseProgram(shader.id)
}

// GetUniform will return the ID for a GLSL uniform
func (shader *Shader) GetUniform(name string) int32 {
	if id, ok := shader.uniforms[name]; ok {
		return id
	}
	return 0
}

// GetShaderByName retrieves a loaded shader by its name
func GetShaderByName(name string) (*Shader, error) {
	if shader, ok := shaderList[name]; ok {
		return &shader, nil
	}
	return nil, errors.New("shader: not found")
}

// New creates a new shader object. The shader can later be activated by the given name with the Use function
func New(name, vertexShaderSource, fragmentShaderSource string, uniforms []string) error {
	if _, found := shaderList[name]; found {
		return fmt.Errorf("Shader already exists: %v", name)
	}
	program, err := newProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		return err
	}
	shader := Shader{}
	shader.id = program
	shader.loaded = true
	shader.vert = vertexShaderSource
	shader.frag = fragmentShaderSource

	gl.UseProgram(shader.id)

	shader.uniforms = map[string]int32{}
	for _, key := range uniforms {
		shader.uniforms[key] = gl.GetUniformLocation(program, gl.Str(key+"\x00"))
	}

	shaderList[name] = shader

	if activeShader := GetActive(); activeShader != nil {
		activeShader.Activate()
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
