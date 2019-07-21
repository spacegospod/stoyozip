package processing

import (
	szio "stoyozip/io"
)

type Compressor struct {
	window []byte // The window of already processed bytes
	lookaheadBuf []byte // The byte sequence yet to be processed
}

// Constructor function
func NewCompressor() *Compressor {
	c := new(Compressor)
	c.window = make([]byte, 0, WINDOW_CAP)
	
	return c
}

func (c *Compressor) Run(is *szio.InputFileStream, os *szio.OutputFileStream) {
	c.lookaheadBuf = is.ReadBytes(LOOKAHEAD_BUFFER_CAP)

	for {
		if len(c.lookaheadBuf) == 0 {
			break
		}

		_, l := c.findLongestMatch()
		
		if l == 0 {
			// no match
			os.WriteBytes([]byte { byte(p), byte(l), c.lookaheadBuf[0] })
			c.slide(is, 1)
		} else {
			// match
			os.WriteBytes([]byte { byte(p), byte(l) })
			c.slide(is, l)
		}
	}
}

// Moves the window and the lookahead buffer n bytes forward
func (c *Compressor) slide(is *szio.InputFileStream, n int) {
	// index to slice the window from
	var i int = n 
	
	// capacity not reached yet
	if len(c.window) < WINDOW_CAP {
		// amount remaining bytes in window
		r := WINDOW_CAP - len(c.window)
		
		if n > r {
			i = n - r
		} else {
			i = 0
		}
	}
	
	// prevent slice out of range
	if i > len(c.window) {
		i = len(c.window)
	}
	
	newWindow := append(c.window[i:], c.lookaheadBuf[:n]...)
	c.window = newWindow
	
	newBuffer := append(c.lookaheadBuf[n:], is.ReadBytes(n)...)
	c.lookaheadBuf = newBuffer
}

// Find the longest matching sequence between the lookahead buffer and
// the window.
// The return value is a pair of pointers (p, l) where
// p is the number of positions backwards into the sliding window and
// l is the number of bytes to read after p.
func (c *Compressor) findLongestMatch() (int, int) {
	if len(c.window) == 0 || len(c.lookaheadBuf) == 0 {
		return 0, 0
	}
	
	var p, l int = 0, 0

	// Matches of less than 3 bytes are not efficient
	for i := 3; i <= len(c.lookaheadBuf); i++ {
		// get the pointer for a matching sequence of length i
		matchIndex := c.testSequence(i)
		
		if matchIndex > -1 {
			// match found, update pointers.
			// will do another loop to check for a match of length i + 1
			// TODO: do not do an entire loop for i + 1. Instead, find the pointers
			// to all matches of length i and on the next loop check only those for i + 1
			p = matchIndex
			l = i
		} else {
			// no match found
			break
		}
	}
	
	return p, l
}

// Tests whether a match with the provided length can be found for the lookahead buffer
func (c *Compressor) testSequence(length int) int {
	if length > len(c.window) {
		return -1
	}

	// todo: don't loop backwards
	for i := len(c.window) - length; i > -1; i-- {
		if c.isSequenceMatch(i, length) {
			return len(c.window) - i
		}
	}
	
	return -1
}

func (c *Compressor) isSequenceMatch(windowStartIndex, lookaheadBufferEndIndex int) bool {	
	for i := 0; i < lookaheadBufferEndIndex; i++ {
		if c.window[i + windowStartIndex] != c.lookaheadBuf[i] {
			return false
		}
	}

	return true
}