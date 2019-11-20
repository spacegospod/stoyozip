package processing

import (
	szio "stoyozip/io"
)

const (
	WINDOW_CAP = 255 // bytes
	LOOKAHEAD_BUFFER_CAP = 255 // bytes
)

type Processor interface {
	Run(r *szio.InputFileStream, w *szio.OutputFileStream)
}

