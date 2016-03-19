package font

import (
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
)

var fontList = map[string]*Font{}

// Font object
type Font struct {
	context *freetype.Context
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
	fontList[file] = s
	return fontList[file], nil
}
