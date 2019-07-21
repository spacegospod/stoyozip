package main

import (
	"os"
	"strings"
	"errors"
	"log"
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
	INPUT_FILE_SHORT = "-i"
	OUTPUT_FILE_SHORT = "-o"
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
				
				if parts[0] == INPUT_FILE || parts[0] == INPUT_FILE_SHORT {
					userInput.inputPath = parts[1]
				} else if parts[0] == OUTPUT_FILE || parts[0] == OUTPUT_FILE_SHORT {
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

func processError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// exclude the actual executable
	args := os.Args[1:]

	userInput, err := parseUserInput(args)
	processError(err)

	p := getProcessor(userInput.command)

	in, err := szio.NewInputFileStream(userInput.inputPath)
	processError(err)
	out, err := szio.NewOutputFileStream(userInput.outputPath)
	processError(err)

	p.Run(in, out)
}