package font

import (
	"image"
	"image/draw"
	"io/ioutil"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var fontList = map[string]*Font{}

// Font object
type Font struct {
	ttf *truetype.Font
}

// New returns a newly created Font object
func New(file string) (*Font, error) {

	if font, found := fontList[file]; found {
		return font, nil
	}
	s := &Font{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	s.ttf, err = truetype.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	fontList[file] = s
	return fontList[file], nil
}

// Draw will draw the font with text
func (font *Font) Draw(text string, size float64) {
	context := freetype.NewContext()
	context.SetFont(font.ttf)

	img := image.NewRGBA(image.Rect(0, 0, 800, 600))
	draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)

	context.SetDst(img)
	context.SetClip(img.Bounds())
	context.SetSrc(image.Black)

	context.SetFontSize(size)
	context.SetDPI(300)

	var tex uint32
	gl.GenTextures(1, &tex)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	context.DrawString(text, freetype.Pt(0, 200))
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(img.Rect.Size().X),
		int32(img.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix))

}
