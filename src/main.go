package main

import (
	"log"
	"os"
	"sz"
)

func main() {
	var args []string = os.Args[1:]
	if len(args) < 3 {
		log.Fatal("Invalid number arguments")
	}
	inputPath := args[1]
	outputPath := args[2]
	if args[0] == "-c" {
		sz.Compress(inputPath, outputPath)
	} else if args[0] == "-x" {
		sz.Decompress(inputPath, outputPath)
	} else {
		log.Fatal("Invalid command")
	}
}
