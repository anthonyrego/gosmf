package camera

import (
	"github.com/anthonyrego/dodge/shader"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var state struct {
	activeCam *Camera
}

func init() {
	state.activeCam = &Camera{}
}

// GetActiveCamera get the current camera in the state
func GetActiveCamera() *Camera {
	return state.activeCam
}

// Camera type for camera making
type Camera struct {
	Position   mgl32.Vec3
	Focus      mgl32.Vec3
	Up         mgl32.Vec3
	projection mgl32.Mat4
	loc        mgl32.Mat4
	zDepth     float32
}

// New creates a new camera
// Set setActive to true if you want to have it start as the render camera.
func New(setActive bool) *Camera {
	cam := Camera{}
	if setActive {
		cam.SetActive()
	}
	return &cam
}

// SetActive will set current camera as the render camera
func (cam *Camera) SetActive() {
	state.activeCam = cam
	if shader := shader.GetActive(); shader != nil {
		gl.UniformMatrix4fv(shader.Projection, 1, false, &cam.projection[0])
		gl.UniformMatrix4fv(shader.Camera, 1, false, &cam.loc[0])
	}
}

// SetOrtho will set the camera to an orthogonal projection
func (cam *Camera) SetOrtho(w int, h int, zDepth int) {
	cam.zDepth = float32(zDepth)
	cam.projection = mgl32.Ortho(0, float32(w), float32(h), 0, -1, cam.zDepth)
	cam.update()
}

// SetPerspective set a perspective projection
func (cam *Camera) SetPerspective(angle float32, w int, h int, zDepth int) {
	cam.projection = mgl32.Perspective(mgl32.DegToRad(angle), float32(w)/float32(h), 0.1, 10.0)
	cam.update()
}

// SetPosition2D will adjust the camera for ortho viewing to specified location
func (cam *Camera) SetPosition2D(x float32, y float32) {
	cam.loc = mgl32.LookAtV(mgl32.Vec3{x, y, cam.zDepth}, mgl32.Vec3{x, y, 0}, mgl32.Vec3{0, 1, 0})
	cam.update()
}

func (cam *Camera) update() {
	if shader := shader.GetActive(); shader != nil && cam == state.activeCam {
		gl.UniformMatrix4fv(shader.Projection, 1, false, &cam.projection[0])
		gl.UniformMatrix4fv(shader.Camera, 1, false, &cam.loc[0])
	}
}
