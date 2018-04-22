// package hha is the highly hackable array.
package hha

import (
	"encoding/base64"
	"fmt"
)

type hha struct {
	size  int
	array []byte
}

// Create will return a new initialized highly-hackable-array.
// It will be of size "n", which is given as input.
// All values start at 0.  You should probably only do this once.
func Create(n int) *hha {
	return &hha{
		size:  n,
		array: make([]byte, n),
	}
}

// CopyWrite changes the bytes in the array starting at the "start" position,
// and will overwrite all values in the array until whichever comes first:
// the end of the array itself, or the end of the input data.
func (h *hha) CopyWrite(start int, data []byte) {
	if start > len(h.array) {
		// out of bounds.
		return
	}
	end := start + len(data)
	if end > len(h.array) {
		// out of bounds.
		return
	}
	h.array = append(h.array[:start], append(data, h.array[end:]...)...)
}

// OverWrite is a another, (possibly better, idk yet), way to edit
// the highly hackable array, starting at position "start".
func (h *hha) OverWrite(start int, data []byte) {
	if start > len(h.array) {
		return
	}
	for i := 0; i < len(data); i++ {
		if (start + i) >= len(h.array) {
			break
		}
		h.array[start+i] = data[i]
	}
}

// Read will return the bytes
func (h *hha) Read(start, stop int) []byte {
	if start < 0 || start > len(h.array) {
		// out of bounds error.
		return []byte{}
	}
	if stop < 0 || stop > len(h.array) {
		// out of bounds error.
		return []byte{}
	}
	if stop > start {
		// allowed.  Start will swap with stop so that it becomes tangible.
		start, stop = stop, start
	}
	return h.array
}

func (h *hha) String() string {
	return fmt.Sprintf("%c", h.array)
}

func (h *hha) Ints() string {
	return fmt.Sprint(h.array)
}

func (h *hha) Length() int {
	return len(h.array)
}

func (h *hha) Base64() string {
	return base64.StdEncoding.EncodeToString(h.array)
}
