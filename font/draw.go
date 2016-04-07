package font

import (
	"image/color"

	"github.com/anthonyrego/gosmf/shader"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Billboard is an object for font rendering
type Billboard struct {
	image     uint32
	vao       uint32
	font      *Font
	text      string
	width     int
	height    int
	rgba      color.Color
	size      float64
	dpi       float64
	texWidth  int
	texHeight int
}

// NewBillboard creates a 2D billboard for rendering
func (font *Font) NewBillboard(text string, width int, height int, scale int, size float64, dpi float64, color color.Color) *Billboard {
	b := &Billboard{}

	b.texWidth = width * scale
	b.texHeight = height * scale

	image := font.createTexture(text, b.texWidth, b.texHeight, size, dpi, color)

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

// Draw will draw the billvboard in the x,y and z
func (billboard *Billboard) Draw(x float32, y float32, z float32) {

	model := mgl32.Translate3D(x, y, z)

	if shader := shader.GetActive(); shader != nil {
		gl.UniformMatrix4fv(shader.Model, 1, false, &model[0])
	}

	gl.BindVertexArray(billboard.vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, billboard.image)

	gl.DrawArrays(gl.TRIANGLES, 0, 1*2*3)
}

// SetText updates the billboard text
func (billboard *Billboard) SetText(text string) {
	if billboard.font != nil && text != billboard.text {
		billboard.font.updateTexture(
			billboard.image,
			text,
			billboard.texWidth,
			billboard.texHeight,
			billboard.size,
			billboard.dpi,
			billboard.rgba)
		billboard.text = text
	}
}

// SetColor updates the color of the text
func (billboard *Billboard) SetColor(color color.Color) {
	if billboard.font != nil && color != billboard.rgba {
		billboard.font.updateTexture(
			billboard.image,
			billboard.text,
			billboard.texWidth,
			billboard.texHeight,
			billboard.size,
			billboard.dpi,
			color)
		billboard.rgba = color
	}
}
