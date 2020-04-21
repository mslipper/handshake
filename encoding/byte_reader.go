package encoding

import "io"

type byteReader struct {
	r io.Reader
}

func (b *byteReader) ReadByte() (byte, error) {
	buf := make([]byte, 1)
	if _, err := b.r.Read(buf); err != nil {
		return 0, err
	}
	return buf[0], nil
}
