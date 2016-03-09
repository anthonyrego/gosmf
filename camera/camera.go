package camera

import (
	"github.com/anthonyrego/dodge/render"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Camera type for camera making
type Camera struct {
	Position   mgl32.Vec3
	Focus      mgl32.Vec3
	Up         mgl32.Vec3
	projection mgl32.Mat4
	zDepth     float32
}

// New creates a new camera
func New() *Camera {
	return &Camera{}
}

// SetOrtho will set the camera to an orthogonal projection
func (cam *Camera) SetOrtho(w int, h int, zDepth int) {
	cam.zDepth = float32(zDepth)
	cam.projection = mgl32.Ortho(0, float32(w), float32(h), 0, -1, cam.zDepth)

	if shader := render.GetCurrentShader(); shader != nil {
		gl.UniformMatrix4fv(shader.Projection, 1, false, &cam.projection[0])
	}
}

// SetPerspective set a perspective projection
func (cam *Camera) SetPerspective(angle float32, w int, h int, zDepth int) {
	cam.projection = mgl32.Perspective(mgl32.DegToRad(angle), float32(w)/float32(h), 0.1, 10.0)
	if shader := render.GetCurrentShader(); shader != nil {
		gl.UniformMatrix4fv(shader.Projection, 1, false, &cam.projection[0])
	}
}

// SetPosition2D will adjust the camera for ortho viewing to specified location
func (cam *Camera) SetPosition2D(x float32, y float32) {
	camMat := mgl32.LookAtV(mgl32.Vec3{x, y, cam.zDepth}, mgl32.Vec3{x, y, 0}, mgl32.Vec3{0, 1, 0})
	if shader := render.GetCurrentShader(); shader != nil {
		gl.UniformMatrix4fv(shader.Camera, 1, false, &camMat[0])
	}
}
