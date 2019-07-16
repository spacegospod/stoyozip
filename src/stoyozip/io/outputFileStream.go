package io

import (
	"os"
)

type OutputFileStream struct {
	file *os.File
}

// constructor
func NewOutputFileStream(path string) (*OutputFileStream, error) {
	s := new(OutputFileStream)
	f, err := os.Create(path)
	
	if err != nil {
		return nil, err
	}
	
	s.file = f
	
	return s, nil
}

func (s *OutputFileStream) WriteBytes(b []byte) {
	s.file.Write(b)
}

