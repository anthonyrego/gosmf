package texture

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var textureList = map[string]*Texture{}

// Texture object
type Texture struct {
	id     uint32
	Width  int
	Height int
}

// New returns a newly created Texture
func New(file string) (*Texture, error) {

	if tex, found := textureList[file]; found {
		return tex, nil
	}
	s := &Texture{}
	err := s.create(file)
	if err != nil {
		log.Fatalln("failed to create texture:", err)
		return nil, err
	}
	textureList[file] = s
	return textureList[file], nil
}

func (t *Texture) create(file string) error {
	imgFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var tex uint32
	gl.GenTextures(1, &tex)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	t.id = tex
	t.Width = rgba.Rect.Size().X
	t.Height = rgba.Rect.Size().Y
	return nil
}

// Bind will set the current texture to the active texture
func (t *Texture) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}
