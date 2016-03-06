package render

import (
	"github.com/anthonyrego/dodge/window"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var state struct {
	shader *Shader
}

// Setup2DProjection will get the projection ready for viewing 2D content
func Setup2DProjection(shader *Shader, screen *window.Screen) {
	project := mgl32.Ortho(0, float32(screen.Width), float32(screen.Height), 0, -1, 1000)
	gl.UniformMatrix4fv(shader.projection, 1, false, &project[0])

	// lets move this out to a camera package or something
	cam := mgl32.LookAtV(mgl32.Vec3{0, 0, 1000}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	gl.UniformMatrix4fv(shader.camera, 1, false, &cam[0])
}
