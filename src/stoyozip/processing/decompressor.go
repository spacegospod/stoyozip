package processing

import (
	szio "stoyozip/io"
)

type Decompressor struct {
	window []byte
}

func NewDecompressor() *Decompressor {
	d := new(Decompressor)
	d.window = make([]byte, 0, WINDOW_CAP)

	return d
}

func (d *Decompressor) Run(is *szio.InputFileStream, os *szio.OutputFileStream) {
	for {
		if is.IsEmpty() {
			return
		}
		decodedSequence := d.getNextSequence(is)
		d.slide(decodedSequence)
		os.WriteBytes(decodedSequence)
	}
}

func (d *Decompressor) getNextSequence(is *szio.InputFileStream) []byte {
	pointer := is.ReadBytes(2)

	if pointer[0] == 0 && pointer[1] == 0 {
		return is.ReadBytes(1)
	} else {
		readStart := len(d.window) - int(pointer[0])
		readEnd := readStart + int(pointer[1])

		return d.window[readStart:readEnd]
	}
}

func (d *Decompressor) slide(sequence []byte) {
	// index to slice the window from
	var i int = len(sequence)

	// capacity not reached yet
	if len(d.window) < WINDOW_CAP {
		// amount remaining bytes in window
		r := WINDOW_CAP - len(d.window)

		if len(sequence) > r {
			i = len(sequence) - r
		} else {
			i = 0
		}
	}

	// prevent slice out of range
	if i > len(d.window) {
		i = len(d.window)
	}

	newWindow := append(d.window[i:], sequence...)
	d.window = newWindow
}