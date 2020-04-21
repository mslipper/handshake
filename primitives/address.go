package primitives

import (
	"errors"
	"handshake/encoding"
	"io"
)

type Address struct {
	Version uint8
	Hash    []byte
}

func (a *Address) Encode(w io.Writer) error {
	if err := encoding.WriteUint8(w, a.Version); err != nil {
		return err
	}
	if err := encoding.WriteUint8(w, uint8(len(a.Hash))); err != nil {
		return err
	}
	if _, err := w.Write(a.Hash); err != nil {
		return err
	}
	return nil
}

func (a *Address) Decode(r io.Reader) error {
	version, err := encoding.ReadUint8(r)
	if err != nil {
		return err
	}
	if version > 31 {
		return errors.New("invalid address version")
	}
	size, err := encoding.ReadUint8(r)
	if err != nil {
		return err
	}
	if size < 2 || size > 40 {
		return errors.New("invalid address length")
	}
	hash := make([]byte, size)
	if _, err := io.ReadFull(r, hash); err != nil {
		return err
	}
	a.Version = version
	a.Hash = hash
	return nil
}
