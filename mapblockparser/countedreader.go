package mapblockparser

import (
	"bytes"
)

type CountedReader struct {
	Reader *bytes.Reader
	Count  int
}

func (r *CountedReader) Read(p []byte) (int, error) {
	i, err := r.Reader.Read(p)
	r.Count += i
	return i, err
}

func (r *CountedReader) ReadByte() (byte, error) {
	i, err := r.Reader.ReadByte()
	r.Count++
	return i, err
}
