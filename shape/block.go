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
func NewBlock() (*Block, error) {
	block := &Block{}
	err := block.create()
	if err != nil {
		log.Fatalln("failed to create block shape:", err)
		return nil, err
	}
	return block, nil
}

func (block *Block) create() error {
	var vao uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	shapeVertices := []float32{
		1, 1, 0.0,
		0.0, 0.0, 0.0,
		0.0, 1, 0.0,
		1, 1, 0.0,
		0.0, 0.0, 0.0,
		1, 0.0, 0.0,
	}

	gl.BufferData(gl.ARRAY_BUFFER, len(shapeVertices)*4, gl.Ptr(shapeVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(0)
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))

	block.vao = vao

	return nil
}

type BlockParams struct {
	X, Y, Z                         float32
	RotationX, RotationY, RotationZ float32
	Width, Height                   float32
	R, G, B, A                      float32
}

// Draw will draw the block
func (block *Block) Draw(params BlockParams) {

	model := mgl32.Translate3D(params.X, params.Y, params.Z)
	if params.RotationX != 0 || params.RotationY != 0 || params.RotationZ != 0 {
		model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(params.RotationX), mgl32.Vec3{1, 0, 0}))
		model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(params.RotationY), mgl32.Vec3{0, 1, 0}))
		model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(params.RotationZ), mgl32.Vec3{0, 0, 1}))
	}
	model = model.Mul4(mgl32.Scale3D(params.Width, params.Height, 1))

	if shader := shader.GetActive(); shader != nil {
		gl.UniformMatrix4fv(shader.GetUniform("model"), 1, false, &model[0])
		gl.Uniform4f(shader.GetUniform("color"), params.R, params.G, params.B, params.A)
	}

	gl.BindVertexArray(block.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

// Destroy frees up the vertex array
func (block *Block) Destroy() {
	gl.DeleteVertexArrays(1, &block.vao)
}
