package primitives

import (
	"errors"
	"github.com/mslipper/handshake/encoding"
	"golang.org/x/crypto/blake2b"
)

func HashName(name string) []byte {
	h, _ := blake2b.New256(nil)
	h.Write([]byte(name))
	return h.Sum(nil)
}

func CreateBlind(value uint64, nonce []byte) ([]byte, error) {
	if len(nonce) != 32 {
		return nil, errors.New("nonce must be 32 bytes long")
	}

	h, _ := blake2b.New256(nil)
	_ = encoding.WriteUint64(h, value)
	h.Write(nonce)
	return h.Sum(nil), nil
}
