package sprite

import (
	"log"

	"github.com/anthonyrego/gosmf/shader"
	"github.com/anthonyrego/gosmf/texture"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var spriteList = map[string]*Sprite{}

// Sprite is used for rendering textures in a 2D ortho way
type Sprite struct {
	image *texture.Texture
	vao   uint32
}

// New creates and returns a new Sprite object
func New(file string, width int, height int) (*Sprite, error) {
	if sprite, found := spriteList[file]; found {
		return sprite, nil
	}
	s := &Sprite{}
	err := s.create(file, width, height, 0, 0)
	if err != nil {
		log.Fatalln("failed to create sprite:", err)
		return nil, err
	}
	spriteList[file] = s
	return spriteList[file], nil
}

// NewSheet creates and returns a new Sprite object that has multiple frames based on the width and height
func NewSheet(file string, width int, height int, frames int, framesPerLine int) (*Sprite, error) {
	if sprite, found := spriteList[file]; found {
		return sprite, nil
	}
	s := &Sprite{}
	err := s.create(file, width, height, frames, framesPerLine)
	if err != nil {
		log.Fatalln("failed to create sprite:", err)
		return nil, err
	}
	spriteList[file] = s
	return spriteList[file], nil
}

func (sprite *Sprite) create(file string, width int, height int, frames int, framesPerLine int) error {
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

	frameWidth := float32(width) / float32(image.Width)
	frameHeight := float32(height) / float32(image.Height)

	var spriteVertices []float32

	if frames == 0 {
		spriteVertices = []float32{
			w, h, 0.0, 1.0, 1.0,
			0.0, 0.0, 0.0, 0.0, 0.0,
			0.0, h, 0.0, 0.0, 1.0,

			w, h, 0.0, 1.0, 1.0,
			0.0, 0.0, 0.0, 0.0, 0.0,
			w, 0.0, 0.0, 1.0, 0.0,
		}
	} else {
		for i := 0; i < frames; i++ {
			frameX := float32(i%framesPerLine) * frameWidth
			frameY := float32(i/framesPerLine) * frameHeight
			frameW := frameWidth + frameX
			frameH := frameHeight + frameY

			spriteVertices = append(spriteVertices,
				[]float32{
					w, h, 0.0, frameW, frameH,
					0.0, 0.0, 0.0, frameX, frameY,
					0.0, h, 0.0, frameX, frameH,

					w, h, 0.0, frameW, frameH,
					0.0, 0.0, 0.0, frameX, frameY,
					w, 0.0, 0.0, frameW, frameY,
				}...)
		}
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
func (sprite *Sprite) Draw(x float32, y float32, z float32) {

	model := mgl32.Translate3D(x, y, z)
	// remember this is in radians!
	// model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(90), mgl32.Vec3{0, 0, 1}))
	if shader := shader.GetActive(); shader != nil {
		gl.UniformMatrix4fv(shader.Model, 1, false, &model[0])
	}

	gl.BindVertexArray(sprite.vao)

	sprite.image.Bind()

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

// DrawFrame will draw the sprite in the x,y and z with the specified frame from a spritesheet
func (sprite *Sprite) DrawFrame(x float32, y float32, z float32, frame int) {

	model := mgl32.Translate3D(x, y, z)
	// remember this is in radians!
	// model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(90), mgl32.Vec3{0, 0, 1}))
	if shader := shader.GetActive(); shader != nil {
		gl.UniformMatrix4fv(shader.Model, 1, false, &model[0])
	}

	gl.BindVertexArray(sprite.vao)

	sprite.image.Bind()

	gl.DrawArrays(gl.TRIANGLES, int32(frame*6), 6)
}
