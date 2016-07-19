package audio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
)

func LoadWav(file string) *Sound {
	var thirtytwo uint32
	var sixteen uint16

	w := &Sound{}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return w
	}

	r := bytes.NewReader(b)

	// Ignore some data that we do not need
	binary.Read(r, binary.BigEndian, &thirtytwo)
	binary.Read(r, binary.LittleEndian, &thirtytwo)
	binary.Read(r, binary.BigEndian, &thirtytwo)
	binary.Read(r, binary.BigEndian, &thirtytwo)
	binary.Read(r, binary.LittleEndian, &thirtytwo)
	binary.Read(r, binary.LittleEndian, &sixteen)

	if err := binary.Read(r, binary.LittleEndian, &w.channels); err != nil {
		return w
	}

	if err := binary.Read(r, binary.LittleEndian, &w.frequency); err != nil {
		return w
	}

	// Ignore some more data that we do not need
	binary.Read(r, binary.LittleEndian, &thirtytwo)
	binary.Read(r, binary.LittleEndian, &sixteen)
	binary.Read(r, binary.LittleEndian, &sixteen)
	binary.Read(r, binary.BigEndian, &thirtytwo)

	if err := binary.Read(r, binary.LittleEndian, &w.size); err != nil {
		return w
	}

	w.data = make([]byte, w.size)

	_, err = io.ReadFull(r, w.data)

	if err != nil {
		fmt.Println(err)
	}
	w.attachSoundData()

	return w
}
