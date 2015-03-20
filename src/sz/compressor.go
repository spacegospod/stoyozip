package sz

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
)

func Compress(inputPath, outputPath string) {
	// TODO: read in chunks
	fileBytes, _ := ioutil.ReadFile(inputPath)
	var pos int = 0

	iom := newIomanager(outputPath)

	for {
		if pos >= len(fileBytes) {
			break
		}
		p, l := findLongestChunk(window, fileBytes[pos:])
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, []uint8{
			uint8(p),
			uint8(l),
		})

		if l == 0 {
			// no match
			pushToWindow([]byte{fileBytes[pos]})
			pos += 1
			binary.Write(buf, binary.LittleEndian, fileBytes[pos-1])
		} else {
			pushToWindow(fileBytes[pos:(pos + l)])
			pos += l
		}

		iom.writeChunk(buf.Bytes())
	}

	// clear leftover buffered data
	iom.flush()
}
