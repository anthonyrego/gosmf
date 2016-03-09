package render

import (
	"github.com/anthonyrego/dodge/texture"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"log"
)

var spriteList = map[string]*Sprite{}

// Sprite is used for rendering textures in a 2D ortho way
type Sprite struct {
	image *texture.Texture
	vao   uint32
}

// NewSprite returns a newly created Sprite
func NewSprite(file string, width int, height int) (*Sprite, error) {
	if sprite, found := spriteList[file]; found {
		return sprite, nil
	}
	s := &Sprite{}
	err := s.create(file, width, height)
	if err != nil {
		log.Fatalln("failed to create sprite:", err)
		return nil, err
	}
	spriteList[file] = s
	return spriteList[file], nil
}

func (sprite *Sprite) create(file string, width int, height int) error {
	image, err := texture.New(file)
	if err != nil {
		return err
	}

	var vao uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	w := float32(width)
	h := float32(height)

	spriteVertices := []float32{
		w, h, 0.0, 1.0, 1.0,
		0.0, 0.0, 0.0, 0.0, 0.0,
		0.0, h, 0.0, 0.0, 1.0,

		w, h, 0.0, 1.0, 1.0,
		0.0, 0.0, 0.0, 0.0, 0.0,
		w, 0.0, 0.0, 1.0, 0.0,
	}

	gl.BufferData(gl.ARRAY_BUFFER, len(spriteVertices)*4, gl.Ptr(spriteVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(0)
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(1)
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	sprite.vao = vao
	sprite.image = image

	return nil
}

// Draw will draw the sprite in the x,y and z
func (sprite *Sprite) Draw(x int, y int, z int) {

	model := mgl32.Translate3D(float32(x), float32(y), float32(z))
	// remember this is in radians!
	// model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(90), mgl32.Vec3{0, 0, 1}))
	if shader := GetCurrentShader(); shader != nil {
		gl.UniformMatrix4fv(shader.Model, 1, false, &model[0])
	}

	gl.BindVertexArray(sprite.vao)

	sprite.image.Bind()

	gl.DrawArrays(gl.TRIANGLES, 0, 1*2*3)
}
