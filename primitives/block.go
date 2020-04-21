package primitives

import (
	"handshake/encoding"
	"io"
)

const (
	NonceSize = 24
	MaskSize  = 32
)

type Block struct {
	Nonce        uint32
	Time         uint64
	Hash         [32]byte
	TreeRoot     [32]byte
	ExtraNonce   [NonceSize]byte
	ReservedRoot [32]byte
	WitnessRoot  [32]byte
	MerkleRoot   [32]byte
	Version      uint32
	Bits         uint32
	Mask         [MaskSize]byte
	Transactions []*Transaction
}

func (b *Block) Encode(w io.Writer) error {
	if err := encoding.WriteUint32(w, b.Nonce); err != nil {
		return err
	}
	if err := encoding.WriteUint64(w, b.Time); err != nil {
		return err
	}
	if _, err := w.Write(b.Hash[:]); err != nil {
		return err
	}
	if _, err := w.Write(b.TreeRoot[:]); err != nil {
		return err
	}
	if _, err := w.Write(b.ExtraNonce[:]); err != nil {
		return err
	}
	if _, err := w.Write(b.ReservedRoot[:]); err != nil {
		return err
	}
	if _, err := w.Write(b.WitnessRoot[:]); err != nil {
		return err
	}
	if _, err := w.Write(b.MerkleRoot[:]); err != nil {
		return err
	}
	if err := encoding.WriteUint32(w, b.Version); err != nil {
		return err
	}
	if err := encoding.WriteUint32(w, b.Bits); err != nil {
		return err
	}
	if _, err := w.Write(b.Mask[:]); err != nil {
		return err
	}
	if err := encoding.WriteVarint(w, uint64(len(b.Transactions))); err != nil {
		return err
	}
	for _, tx := range b.Transactions {
		if err := tx.Encode(w); err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) Decode(r io.Reader) error {
	nonce, err := encoding.ReadUint32(r)
	if err != nil {
		return err
	}
	ts, err := encoding.ReadUint64(r)
	if err != nil {
		return err
	}
	var hash [32]byte
	if _, err := r.Read(hash[:]); err != nil {
		return err
	}
	var treeRoot [32]byte
	if _, err := r.Read(treeRoot[:]); err != nil {
		return err
	}
	var extraNonce [NonceSize]byte
	if _, err := r.Read(extraNonce[:]); err != nil {
		return err
	}
	var reservedRoot [32]byte
	if _, err := r.Read(reservedRoot[:]); err != nil {
		return err
	}
	var witnessRoot [32]byte
	if _, err := r.Read(witnessRoot[:]); err != nil {
		return err
	}
	var merkleRoot [32]byte
	if _, err := r.Read(merkleRoot[:]); err != nil {
		return err
	}
	version, err := encoding.ReadUint32(r)
	if err != nil {
		return err
	}
	bits, err := encoding.ReadUint32(r)
	if err != nil {
		return err
	}
	var mask [MaskSize]byte
	if _, err := r.Read(mask[:]); err != nil {
		return err
	}
	txCount, err := encoding.ReadVarint(r)
	if err != nil {
		return err
	}
	var txs []*Transaction
	for i := 0; i < int(txCount); i++ {
		tx := new(Transaction)
		if err := tx.Decode(r); err != nil {
			return err
		}
		txs = append(txs, tx)
	}
	b.Nonce = nonce
	b.Time = ts
	b.Hash = hash
	b.TreeRoot = treeRoot
	b.ExtraNonce = extraNonce
	b.ReservedRoot = reservedRoot
	b.WitnessRoot = witnessRoot
	b.MerkleRoot = merkleRoot
	b.Version = version
	b.Bits = bits
	b.Mask = mask
	b.Transactions = txs
	return nil
}
