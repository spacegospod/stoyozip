package sz

import (
	"io/ioutil"
)

func Decompress(inputPath, outputPath string) {
	fileBytes, _ := ioutil.ReadFile(inputPath)
	var pos int = 0

	iom := newIomanager(outputPath)

	for {
		if pos >= len(fileBytes)-1 {
			break
		}

		var p, l uint8
		p = uint8(fileBytes[pos])
		l = uint8(fileBytes[pos+1])
		if l == 0 {
			var b byte = fileBytes[pos+2]
			iom.writeChunk([]byte{b})
			pushToWindow([]byte{b})
			pos += 3
		} else {
			// invert
			p = uint8(len(window)) - p
			var b_sec []byte = window[p:(p + l)]
			pushToWindow(b_sec)
			iom.writeChunk(b_sec)
			pos += 2
		}
	}

	// clear leftover buffered data
	iom.flush()
}
