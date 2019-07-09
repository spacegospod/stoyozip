package main

import (
	szio "stoyozip/io"
	"stoyozip/processing"
)

func main() {
	r := szio.NewInputFileStream("E:\\git\\stoyozip\\input.txt")
	w := szio.NewOutputFileStream("E:\\git\\stoyozip\\output.txt")
	
	c := processing.NewCompressor()
	c.Run(r, w)
}