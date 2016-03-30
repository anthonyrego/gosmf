package font

import (
	"image/color"

	"github.com/anthonyrego/dodge/shader"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Billboard is an object for font rendering
type Billboard struct {
	image  uint32
	vao    uint32
	font   *Font
	text   string
	width  int
	height int
	rgba   color.Color
	size   float64
	dpi    float64
}

// NewBillboard creates a 2D billboard for rendering
func (font *Font) NewBillboard(text string, width int, height int, size float64, dpi float64, color color.Color) *Billboard {
	b := &Billboard{}

	image := font.createTexture(text, width, height, size, dpi, color)

	b.width = width
	b.height = height
	b.size = size
	b.dpi = dpi
	b.text = text
	b.font = font
	b.rgba = color

	var vao uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	w := float32(width)
	h := float32(height)

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

// Draw will draw the sprite in the x,y and z
func (billboard *Billboard) Draw(x float32, y float32, z float32) {

	model := mgl32.Translate3D(x, y, z)
	// remember this is in radians!
	// model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(90), mgl32.Vec3{0, 0, 1}))
	if shader := shader.GetActive(); shader != nil {
		gl.UniformMatrix4fv(shader.Model, 1, false, &model[0])
	}

	gl.BindVertexArray(billboard.vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, billboard.image)

	gl.DrawArrays(gl.TRIANGLES, 0, 1*2*3)
}

// UpdateText billboard text
func (billboard *Billboard) UpdateText(text string) {
	if billboard.font != nil && text != billboard.text {
		billboard.font.updateTexture(
			billboard.image,
			text,
			billboard.width,
			billboard.height,
			billboard.size,
			billboard.dpi,
			billboard.rgba)
		billboard.text = text
	}
}
