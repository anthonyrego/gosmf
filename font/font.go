package font

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"

	"golang.org/x/image/math/fixed"

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

func (font *Font) createTexture(text string, width int, height int, size float64, dpi float64) (uint32, int, int) {
	context := freetype.NewContext()
	context.SetFont(font.ttf)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.RGBA{255, 255, 255, 0}), image.ZP, draw.Src)

	context.SetDst(img)
	context.SetClip(img.Bounds())
	context.SetSrc(image.NewUniform(color.RGBA{255, 255, 255, 255}))
	context.SetFontSize(size)
	context.SetDPI(dpi)
	pixelBounds, _ := context.DrawString(text, freetype.Pt(0, height/2))

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

	return tex, int26_6Ceiling(pixelBounds.X + 0x3f), int26_6Ceiling(pixelBounds.Y + 0x3f)
}

func (font *Font) updateTexture(texture uint32, text string, width int, height int, size float64, dpi float64) (int, int) {
	context := freetype.NewContext()
	context.SetFont(font.ttf)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.RGBA{255, 255, 255, 0}), image.ZP, draw.Src)

	context.SetDst(img)
	context.SetClip(img.Bounds())
	context.SetSrc(image.NewUniform(color.RGBA{255, 255, 255, 255}))
	context.SetFontSize(size)
	context.SetDPI(dpi)
	pixelBounds, _ := context.DrawString(text, freetype.Pt(0, height/2))

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexSubImage2D(
		gl.TEXTURE_2D,
		0,
		0,
		0,
		int32(img.Rect.Size().X),
		int32(img.Rect.Size().Y),
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix))

	return int26_6Ceiling(pixelBounds.X + 0x3f), int26_6Ceiling(pixelBounds.Y + 0x3f)
}

func int26_6Ceiling(x fixed.Int26_6) int {
	return int((x + 0x3f) >> 6)
}
