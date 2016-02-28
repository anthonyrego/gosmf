package render

import (
	"github.com/anthonyrego/dodge/texture"
	"github.com/go-gl/gl/v4.1-core/gl"
)

var vao uint32

// Init2DSprite is a test function
func Init2DSprite(shader *Shader) {

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(spriteVertices)*4, gl.Ptr(spriteVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(shader.id, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(shader.id, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

}

// DrawSprite will draw your Sprites, duh
func DrawSprite(tex *texture.Texture, w int, h int, x int, y int) {

	gl.BindVertexArray(vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex.ID)

	gl.DrawArrays(gl.TRIANGLES, 0, 1*2*3)
}

var spriteVertices = []float32{
	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,

	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 1.0, 0.0,
}
