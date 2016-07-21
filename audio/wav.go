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

	w := NewSound(file)
	if w.buffer != 0 {
		return w
	}
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

	if err := binary.Read(r, binary.LittleEndian, &w.Channels); err != nil {
		return w
	}

	if err := binary.Read(r, binary.LittleEndian, &w.Frequency); err != nil {
		return w
	}

	// Ignore some more data that we do not need
	binary.Read(r, binary.LittleEndian, &thirtytwo)
	binary.Read(r, binary.LittleEndian, &sixteen)
	if err := binary.Read(r, binary.LittleEndian, &w.BitsPerSample); err != nil {
		return w
	}
	binary.Read(r, binary.BigEndian, &thirtytwo)

	if err := binary.Read(r, binary.LittleEndian, &w.Size); err != nil {
		return w
	}

	w.Data = make([]byte, w.Size)

	_, err = io.ReadFull(r, w.Data)

	if err != nil {
		fmt.Println(err)
	}
	w.LoadPCMData()

	return w
}
