package io

import (
	"io/ioutil"
)

type InputFileStream struct {
	buffer []byte
	bufferIndex int
}

// constructor
func NewInputFileStream(path string) *InputFileStream {
	// todo handle missing file and nil path
	s := new(InputFileStream)
	fileBytes, _ := ioutil.ReadFile(path)
	s.buffer = fileBytes
	s.bufferIndex = 0
	
	return s
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