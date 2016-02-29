package render

import (
	"github.com/anthonyrego/dodge/texture"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// DrawSprite will draw your Sprites, duh
func DrawSprite(tex *texture.Texture, w int, h int, x int, y int) {

	gl.BindVertexArray(tex.VertexArrayObject)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex.ID)

	gl.DrawArrays(gl.TRIANGLES, 0, 1*2*3)
}
