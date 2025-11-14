package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int       // position in the chain
	Timestamp    time.Time // when the block was created
	Data         string    // payload or transaction
	Transactions []Transaction
	PrevHash     string // hash of the previous block
	Hash         string // current block hash
	Nonce        int    //

	//    Nonce        int

}

// a method to calculate block's hash

// a method to calculate block's hash
func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp.String() + b.PrevHash + string(b.Nonce)
	for _, tx := range b.Transactions {
		record += tx.Sender + tx.Receiver + string(tx.Amount) + tx.Signature
	}
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

/*
func (b * )



*/

// Note: SHA‑256 is a cryptographic hash function that takes any input
// and produces a fixed 256‑bit (64‑character) output.
// Note: SHA-256 is a cryptographic hash function that takes any input and produces a fixed 256-bit (64 characters)
