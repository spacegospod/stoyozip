/*
	Algorithm taken from https://msdn.microsoft.com/en-us/library/ee916854.aspx
*/
package sz

const WINDOW_SIZE int = 1024

var window []byte = make([]byte, 0, WINDOW_SIZE)

func findLongestChunk(window, buf []byte) (int, int) {
	var p, l int = 0, 0
	if len(window) == 0 || len(buf) == 0 {
		return 0, 0
	}

	for {
		var chunk_len int = l + 1
		if chunk_len > len(buf) {
			break
		}
		begin_index := findChunk(window, buf[:chunk_len])
		if begin_index == -1 {
			break
		}
		p = begin_index
		l = chunk_len
	}

	if l == 0 {
		return 0, 0
	} else {
		return (len(window) - p), l
	}
}

func findChunk(window, chunk []byte) int {
	var p int = -1
	for i, _ := range window[:(len(window)-len(chunk))+1] {
		var found bool = true
		for j := 0; j < len(chunk); j++ {
			if window[i+j] != chunk[j] {
				found = false
			}
		}

		if found {
			p = i
		}
	}
	return p
}

func pushToWindow(bytes []byte) {
	window = append(window, bytes...)

	var excess = len(window) - WINDOW_SIZE
	if excess > 0 {
		window = window[excess:]
	}
}
