package lib

import (
	"io"
)

const chunkSize = 16

type Reader struct {
	input io.Reader
}

func NewReader(input io.Reader) *Reader {
	return &Reader{input: input}
}

func (r *Reader) ReadChunk() (int64, []byte, error) {
	buffer := make([]byte, chunkSize)
	n, err := r.input.Read(buffer)
	if n == 0 {
		return 0, nil, err
	}
	return int64(n), buffer[:n], err
}
