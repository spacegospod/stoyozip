package main

import (
	"os"
	"log"
	"strings"
	"errors"
	szio "stoyozip/io"
	"stoyozip/processing"
)

const (
	// commands
	COMPRESS = "-c"
	DECOMPRESS = "-x"
	
	// parameters
	INPUT_FILE = "-inputFile"
	OUTPUT_FILE = "-outputFile"
)

type UserInput struct {
	command, inputPath, outputPath string
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

func parseUserInput(args []string) (UserInput, error) {
	var userInput UserInput
	var error error = nil

	if len(args) < 3 {
		error = errors.New("Invalid number arguments")
	}
	
	if error == nil {
		for _, v := range args {
			if v == COMPRESS || v == DECOMPRESS {
				userInput.command = v
			} else {
				parts := strings.Split(v, "=")
				
				if parts[0] == INPUT_FILE {
					userInput.inputPath = parts[1]
				} else if parts[0] == OUTPUT_FILE {
					userInput.outputPath = parts[1]
				}
			}
		}
	}
	
	error = validateUserInput(userInput)
	
	return userInput, error
}

func validateUserInput(userInput UserInput) error {
	if len(userInput.command) == 0 {
		return errors.New("Invalid command")
	}
	
	if len(userInput.inputPath) == 0 {
		return errors.New("Invalid input file")
	}
	
	if len(userInput.outputPath) == 0 {
		return errors.New("Invalid output file")
	}
	
	return nil
}

func main() {
	// exclude the actual executable
	args := os.Args[1:]

	userInput, err := parseUserInput(args)
	
	if err != nil {
		log.Fatal(err)
	}

	p := getProcessor(userInput.command)

	in := szio.NewInputFileStream(userInput.inputPath)
	out := szio.NewOutputFileStream(userInput.outputPath)

	p.Run(in, out)
}