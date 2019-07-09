package processing

import (
	szio "stoyozip/io"
)

type Processor interface {
	Run(r *szio.InputFileStream, w *szio.OutputFileStream)
}

