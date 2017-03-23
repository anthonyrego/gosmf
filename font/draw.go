package font

import (
	"github.com/anthonyrego/gosmf/shader"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Billboard is an object for font rendering
type Billboard struct {
	Width     int
	Height    int
	image     uint32
	vao       uint32
	font      *Font
	text      string
	size      float64
	dpi       float64
	texWidth  int
	texHeight int
}

type NewBillboardParams struct {
	Text      string
	MaxWidth  int
	MaxHeight int
	Size      float64
	Dpi       float64
}

// NewBillboard creates a 2D billboard for rendering
func (font *Font) NewBillboard(params NewBillboardParams) *Billboard {
	b := &Billboard{}

	b.texWidth = params.MaxWidth
	b.texHeight = params.MaxHeight

	image, renderedWidth, renderedHeight := font.createTexture(params.Text, b.texWidth, b.texHeight, params.Size, params.Dpi)

	b.size = params.Size
	b.dpi = params.Dpi
	b.text = params.Text
	b.font = font

	b.Width = renderedWidth
	b.Height = renderedHeight
	var vao uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	w := float32(params.MaxWidth)
	h := float32(params.MaxHeight)

	billboardVertices := []float32{
		w, h, 0.0, 1.0, 1.0,
		0.0, 0.0, 0.0, 0.0, 0.0,
		0.0, h, 0.0, 0.0, 1.0,

		w, h, 0.0, 1.0, 1.0,
		0.0, 0.0, 0.0, 0.0, 0.0,
		w, 0.0, 0.0, 1.0, 0.0,
	}

	gl.BufferData(gl.ARRAY_BUFFER, len(billboardVertices)*4, gl.Ptr(billboardVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(0)
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(1)
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	b.vao = vao
	b.image = image

	return b
}

type DrawBillboardParams struct {
	X, Y, Z                         float32
	RotationX, RotationY, RotationZ float32
	Scale                           bool
	ScaleX, ScaleY, ScaleZ          float32
	R, G, B, A                      float32
}

// Draw will draw the billboard
func (billboard *Billboard) Draw(params DrawBillboardParams) {

	model := mgl32.Translate3D(params.X, params.Y, params.Z)
	if params.RotationX != 0 || params.RotationY != 0 || params.RotationZ != 0 {
		model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(params.RotationX), mgl32.Vec3{1, 0, 0}))
		model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(params.RotationY), mgl32.Vec3{0, 1, 0}))
		model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(params.RotationZ), mgl32.Vec3{0, 0, 1}))
	}
	if params.Scale {
		model = model.Mul4(mgl32.Scale3D(params.ScaleX, params.ScaleY, params.ScaleZ))
	}
	if shader := shader.GetActive(); shader != nil {
		gl.UniformMatrix4fv(shader.GetUniform("model"), 1, false, &model[0])
		gl.Uniform4f(shader.GetUniform("color"), params.R, params.G, params.B, params.A)
	}

	gl.BindVertexArray(billboard.vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, billboard.image)

	gl.DrawArrays(gl.TRIANGLES, 0, 1*2*3)
}

// SetText updates the billboard text
func (billboard *Billboard) SetText(text string) {
	if billboard.font != nil && text != billboard.text {
		renderedWidth, renderedHeight := billboard.font.updateTexture(
			billboard.image,
			text,
			billboard.texWidth,
			billboard.texHeight,
			billboard.size,
			billboard.dpi)
		billboard.text = text
		billboard.Width = renderedWidth
		billboard.Height = renderedHeight
	}
}

// Destroy frees up the vertex array and texture
func (billboard *Billboard) Destroy() {
	gl.DeleteVertexArrays(1, &billboard.vao)
	gl.DeleteTextures(1, &billboard.image)
}
