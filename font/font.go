package font

import (
	"image"
	"image/draw"
	"io/ioutil"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/golang/freetype"
)

var fontList = map[string]*Font{}

// Font object
type Font struct {
	context *freetype.Context
	image   *image.RGBA
	texID   uint32
}

// New returns a newly created Font object
func New(file string) (*Font, error) {

	if font, found := fontList[file]; found {
		return font, nil
	}
	s := &Font{}
	s.context = freetype.NewContext()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	ttf, err := freetype.ParseFont(data)
	if err != nil {
		log.Fatal(err)
	}
	s.context.SetFont(ttf)

	s.image = image.NewRGBA(image.Rect(0, 0, 800, 600))
	draw.Draw(s.image, s.image.Bounds(), image.Transparent, image.ZP, draw.Src)

	s.context.SetDst(s.image)
	s.context.SetClip(s.image.Bounds())
	s.context.SetSrc(image.Black)

	s.context.SetFontSize(12)
	s.context.SetDPI(300)
	//s.context.DrawString("Wham: Make it Big", freetype.Pt(100, 100))
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
		int32(s.image.Rect.Size().X),
		int32(s.image.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(s.image.Pix))

	s.texID = tex

	fontList[file] = s
	return fontList[file], nil
}

// Draw will draw the font with text in the x,y and z
func (font *Font) Draw(text string, x float32, y float32, z float32) {

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, font.texID)
	font.context.DrawString(text, freetype.Pt(0, 200))

}
