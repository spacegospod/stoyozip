package main

import (
	"os"
	"log"
	szio "stoyozip/io"
	"stoyozip/processing"
)

const (
	// commands
	COMPRESS = "-c"
	DECOMPRESS = "-x"
)

func main() {
	// TODO: add proper error handling
	args := os.Args[1:]

	if len(args) < 3 {
		log.Fatal("Invalid number arguments")
	}

	command := args[0]
	inputPath := args[1]
	outputPath := args[2]

	p := getProcessor(command)

	in := szio.NewInputFileStream(inputPath)
	out := szio.NewOutputFileStream(outputPath)

	p.Run(in, out)
}

func getProcessor(command string) processing.Processor {
	if command == COMPRESS {
		return processing.NewCompressor()
	} else if command == DECOMPRESS {
		return processing.NewDecompressor()
	}

	log.Fatal("Invalid command")
	return nil
}