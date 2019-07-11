package processing

import (
	szio "stoyozip/io"
	"log"
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
		log.Println(decodedSequence)
	}
}

func (d *Decompressor) getNextSequence(is *szio.InputFileStream) []byte {
	pointer := is.ReadBytes(2)
	
	if pointer[0] == 0 && pointer[1] == 0 {
		return is.ReadBytes(1)
	} else {
		readIndex := len(d.window) - int(pointer[0])
		readLength := readIndex + int(pointer[1])
		return d.window[readIndex:readLength]
	}
}

func (d *Decompressor) slide(sequence []byte) {
	wi := len(sequence) - 1;
	if wi > len(d.window) {
		wi = 0
	}

	newWindow := append(d.window[wi:], sequence...)
	d.window = newWindow
}