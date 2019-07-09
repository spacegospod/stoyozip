package io

import (
	"os"
	"log"
)

type OutputFileStream struct {
	file *os.File
}

// constructor
func NewOutputFileStream(path string) *OutputFileStream {
	// todo handle nil path
	s := new(OutputFileStream)
	f, _ := os.Create(path)
	
	s.file = f
	
	return s
}

func (s *OutputFileStream) WriteBytes(b []byte) {
	s.file.Write(b)
	log.Println(b)
}

