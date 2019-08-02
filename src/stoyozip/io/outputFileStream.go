package io

import (
	"os"
	"sync"
)

const (
	OUTPUT_BUFFER_SIZE = 256
)

type OutputFileStream struct {
	primaryBuffer []byte
	secondaryBuffer []byte
	file *os.File

	// locks
	swapMutex sync.Mutex
	writeMutex sync.Mutex
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
	s.writeMutex.Lock()
	for {
		if b == nil || len(b) == 0 {
			s.writeMutex.Unlock()
			return
		}

		var numBytesToWrite int = len(b)
		var numBytesAvailable int = cap(s.primaryBuffer) - len(s.primaryBuffer)
		var shouldSwap bool = false
		
		if numBytesToWrite > numBytesAvailable {
			numBytesToWrite = numBytesAvailable
			shouldSwap = true
		}
		
		bytesToWrite := b[:numBytesToWrite];
		s.primaryBuffer = append(s.primaryBuffer, bytesToWrite...)
		b = b[numBytesToWrite:]
		
		if shouldSwap {
			s.swapMutex.Lock()
			// swap buffers
			s.primaryBuffer, s.secondaryBuffer = s.secondaryBuffer, s.primaryBuffer
			// write the secondary to the output file
			go s.writeSecondaryToFile()
		}
	}
}

func (s *OutputFileStream) Flush() {
	s.writeMutex.Lock()
	s.swapMutex.Lock()
	s.file.Write(s.primaryBuffer)
	s.writeMutex.Unlock()
}

func (s *OutputFileStream) writeSecondaryToFile() {
	s.file.Write(s.secondaryBuffer)
	s.secondaryBuffer = s.createEmptyBuffer()
	s.swapMutex.Unlock()
}

func (s *OutputFileStream) createEmptyBuffer() []byte {
	return make([]byte, 0, OUTPUT_BUFFER_SIZE)
}