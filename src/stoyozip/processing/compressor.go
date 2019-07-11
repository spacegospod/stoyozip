package processing

import (
	szio "stoyozip/io"
)

type Compressor struct {
	window []byte
	lookaheadBuf []byte
}

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

		p, l := c.findLongestMatch()
		
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

func (c *Compressor) slide(is *szio.InputFileStream, n int) {
	wi := n - 1;
	if wi > len(c.window) {
		wi = 0
	}
	newWindow := append(c.window[wi:], c.lookaheadBuf[:n]...)
	c.window = newWindow
	
	newBuffer := append(c.lookaheadBuf[n:], is.ReadBytes(n)...)
	c.lookaheadBuf = newBuffer
}

func (c *Compressor) findLongestMatch() (int, int) {
	if len(c.window) == 0 || len(c.lookaheadBuf) == 0 {
		return 0, 0
	}
	
	var p, l int = 0, 0
	
	// todo: i := 3
	// Matches of less than 3 bytes are not efficient
	for i := 1; i < len(c.lookaheadBuf) + 1; i++ {
		testSlice := c.lookaheadBuf[:i]
		matchIndex := c.testSequence(testSlice)
		
		if matchIndex > -1 {
			p = matchIndex
			l = len(testSlice)
		} else {
			break
		}
	}
	
	return p, l
}

func (c *Compressor) testSequence(sequence []byte) int {
	if len(sequence) > len(c.window) {
		return -1
	}

	// todo: don't loop backwards
	for i := len(c.window) - len(sequence); i > -1; i-- {
		if isSequenceMatch(c.window[i:(i + len(sequence))], sequence) {
			return len(c.window) - i
		}
	}
	
	return -1
}

func isSequenceMatch(windowSlice, sequence []byte) bool {
	// asumed to be of equal length
	for i := 0; i < len(windowSlice); i++ {
		if windowSlice[i] != sequence[i] {
			return false
		}
	}

	return true
}