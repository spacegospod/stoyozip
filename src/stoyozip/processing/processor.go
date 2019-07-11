package processing

import (
	szio "stoyozip/io"
)

const (
	WINDOW_CAP = 256 // bytes
	LOOKAHEAD_BUFFER_CAP = 256 // bytes
)

type Processor interface {
	Run(r *szio.InputFileStream, w *szio.OutputFileStream)
}

