package render

var state struct {
	shader *Shader
}

// GetCurrentShader get the current shader in the state
func GetCurrentShader() *Shader {
	return state.shader
}
