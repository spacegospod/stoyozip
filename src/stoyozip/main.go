package main

import (
	szio "stoyozip/io"
	"stoyozip/processing"
)

func main() {
	r := szio.NewInputFileStream("E:\\git\\stoyozip\\output.txt")
	w := szio.NewOutputFileStream("E:\\git\\stoyozip\\restored.txt")
	
	c := processing.NewDecompressor()
	c.Run(r, w)
}