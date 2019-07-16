package io

import (
	"io/ioutil"
)

type InputFileStream struct {
	buffer []byte
	bufferIndex int
}

// constructor
func NewInputFileStream(path string) (*InputFileStream, error) {
	s := new(InputFileStream)
	fileBytes, err := ioutil.ReadFile(path)
	
	if err != nil {
		return nil, err
	}

	s.buffer = fileBytes
	s.bufferIndex = 0
	
	return s, nil
}

func (s *InputFileStream) ReadBytes(n int) []byte {
	var readStart = s.bufferIndex
	var readEnd = readStart + n
	if readEnd > len(s.buffer) {
		readEnd = len(s.buffer)
	}
	var result []byte = s.buffer[readStart:readEnd]
	s.bufferIndex = readEnd
	return result
}

func (s *InputFileStream) IsEmpty() bool {
	return s.bufferIndex >= len(s.buffer)
}