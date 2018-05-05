package camera

import (
	"github.com/anthonyrego/gosmf/shader"
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
	Bounds     mgl32.Vec3
	LookAt     mgl32.Vec3
	Up         mgl32.Vec3
	projection mgl32.Mat4
	viewMatrix mgl32.Mat4
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
	if s := shader.GetActive(); s != nil {
		gl.UniformMatrix4fv(s.GetUniform("projection"), 1, false, &cam.projection[0])
		gl.UniformMatrix4fv(s.GetUniform("camera"), 1, false, &cam.viewMatrix[0])
	}
}

// SetOrtho will set the camera to an orthogonal projection
func (cam *Camera) SetOrtho(w int, h int, zDepth int) {
	cam.zDepth = float32(zDepth)
	cam.Bounds = mgl32.Vec3{float32(w), float32(h), float32(zDepth)}
	cam.projection = mgl32.Ortho(0, float32(w), float32(h), 0, -1, cam.zDepth)
	cam.update()
}

// SetPerspective set a perspective projection
func (cam *Camera) SetPerspective(angle float32, w int, h int, zDepth int) {
	cam.zDepth = float32(zDepth)
	cam.Bounds = mgl32.Vec3{float32(w), float32(h), float32(zDepth)}
	cam.projection = mgl32.Perspective(mgl32.DegToRad(angle), float32(w)/float32(h), 0.1, 10.0)
	cam.update()
}

// SetPosition2D will adjust the camera for ortho viewing to specified location
func (cam *Camera) SetPosition2D(x float32, y float32) {
	cam.Position = mgl32.Vec3{x, y, cam.zDepth}
	cam.viewMatrix = mgl32.LookAtV(mgl32.Vec3{x, y, cam.zDepth}, mgl32.Vec3{x, y, 0}, mgl32.Vec3{0, 1, 0})
	cam.update()
}

// SetViewMatrix will set the lookAt and up vector as well as the position of the camera
func (cam *Camera) SetViewMatrix(position mgl32.Vec3, lookAt mgl32.Vec3, up mgl32.Vec3) {
	cam.Position = position
	cam.LookAt = lookAt
	cam.Up = up
	cam.viewMatrix = mgl32.LookAtV(position, lookAt, up)
	cam.update()
}

// GetViewMatrix will get matrix for the camera
func (cam *Camera) GetViewMatrix() mgl32.Mat4 {
	return cam.viewMatrix
}

func (cam *Camera) update() {
	if s := shader.GetActive(); s != nil && cam == state.activeCam {
		gl.UniformMatrix4fv(s.GetUniform("projection"), 1, false, &cam.projection[0])
		gl.UniformMatrix4fv(s.GetUniform("camera"), 1, false, &cam.viewMatrix[0])
	}
}
