package io

import (
	"os"
	"io"
	"sync"
)

const (
	INPUT_BUFFER_SIZE = 4 * 1024 // 4KB
)

type InputFileStream struct {
	primaryBuffer []byte
	secondaryBuffer []byte
	bufferIndex int
	file *os.File

	// locks
	swapMutex sync.Mutex
}

// constructor
func NewInputFileStream(path string) (*InputFileStream, error) {
	s := new(InputFileStream)

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	s.file = f

	// prepare both buffers
	s.primaryBuffer = s.getNextBuffer()
	s.secondaryBuffer = s.getNextBuffer()

	s.bufferIndex = 0
	
	return s, nil
}

// Reads n bytes from the stream. If less than n bytes are available
// reads as much as possible.
func (s *InputFileStream) ReadBytes(n int) []byte {
	var result []byte = make([]byte, 0, n)
	var remaining = n

	for {
		// done reading or nothing left in stream
		if remaining == 0 || s.IsEmpty() {
			break
		}

		var readStart = s.bufferIndex
		var readEnd = readStart + remaining

		// If there aren't enough bytes in the primary buffer
		// read only what's available.
		if readEnd > len(s.primaryBuffer) {
			readEnd = len(s.primaryBuffer)
		}
		result = append(result, s.primaryBuffer[readStart:readEnd]...)
		s.bufferIndex = readEnd

		var bytesRead = readEnd - readStart
		remaining -= bytesRead

		// Primary buffer depleted, swap with secondary.
		if s.bufferIndex == len(s.primaryBuffer) {
			// Do not allow another swap before the secondary
			// has been refilled.
			s.swapMutex.Lock()
			s.primaryBuffer = s.secondaryBuffer
			s.bufferIndex = 0
			go s.refillSecondaryBuffer()
		}
	}

	return result
}

func (s *InputFileStream) IsEmpty() bool {
	return len(s.primaryBuffer) == 0 && len(s.secondaryBuffer) == 0
}

func (s *InputFileStream) refillSecondaryBuffer() {
	s.secondaryBuffer = s.getNextBuffer()
	s.swapMutex.Unlock()
}

func (s *InputFileStream) getNextBuffer() []byte {
	buffer := make([]byte, INPUT_BUFFER_SIZE, INPUT_BUFFER_SIZE)

	n, err := s.file.Read(buffer)
	
	if n == 0 && err == io.EOF {
		return make([]byte, 0, 0)
	}

	return buffer[:n]
}