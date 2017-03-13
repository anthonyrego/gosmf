package shape

import (
	"log"

	"github.com/anthonyrego/gosmf/shader"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Block is the struct for displaying a colored block
type Block struct {
	vao uint32
}

// New creates and returns a new Block object
func NewBlock(width, height int) (*Block, error) {
	block := &Block{}
	err := block.create(width, height)
	if err != nil {
		log.Fatalln("failed to create block shape:", err)
		return nil, err
	}
	return block, nil
}

func (block *Block) create(width, height int) error {
	var vao uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	w := float32(width)
	h := float32(height)

	shapeVertices := []float32{
		w, h, 0.0,
		0.0, 0.0, 0.0,
		0.0, h, 0.0,
		w, h, 0.0,
		0.0, 0.0, 0.0,
		w, 0.0, 0.0,
	}

	gl.BufferData(gl.ARRAY_BUFFER, len(shapeVertices)*4, gl.Ptr(shapeVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(0)
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))

	block.vao = vao

	return nil
}

// Draw will draw the block in the x,y and z
func (block *Block) Draw(x, y, z, r, g, b, a, scale float32) {

	model := mgl32.Translate3D(x, y, z)
	model = model.Mul4(mgl32.Scale3D(scale, scale, 1))
	// remember this is in radians!
	// model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(90), mgl32.Vec3{0, 0, 1}))
	if shader := shader.GetActive(); shader != nil {
		gl.UniformMatrix4fv(shader.GetUniform("model"), 1, false, &model[0])
		gl.Uniform4f(shader.GetUniform("color"), r, g, b, a)
	}

	gl.BindVertexArray(block.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}
