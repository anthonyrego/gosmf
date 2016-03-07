package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var state struct {
	shader *Shader
}

// Setup2DProjection will get the projection ready for viewing 2D content
func Setup2DProjection(width int, height int) {
	project := mgl32.Ortho(0, float32(width), float32(height), 0, -1, 1000)
	gl.UniformMatrix4fv(state.shader.projection, 1, false, &project[0])
}

// Set2DCamera will adjust the camera for ortho viewing to specified location
func Set2DCamera(x float32, y float32) {
	cam := mgl32.LookAtV(mgl32.Vec3{x, y, 1000}, mgl32.Vec3{x, y, 0}, mgl32.Vec3{0, 1, 0})
	gl.UniformMatrix4fv(state.shader.camera, 1, false, &cam[0])
}
