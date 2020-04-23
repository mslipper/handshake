package primitives

import (
	"golang.org/x/crypto/blake2b"
)

func HashName(name string) []byte {
	h, _ := blake2b.New256(nil)
	h.Write([]byte(name))
	return h.Sum(nil)
}
