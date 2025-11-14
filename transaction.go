// Developed by: Haiylu T.Kibret

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    int
	Signature string
}

// GenerateKeyPair creates a new ECDSA private/public key pair
func GenerateKeyPair() (*ecdsa.PrivateKey, ecdsa.PublicKey) {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return privKey, privKey.PublicKey
}

// SignTransaction signs transaction data with sender’s private key
func (tx *Transaction) SignTransaction(privKey *ecdsa.PrivateKey) {
	data := tx.Sender + tx.Receiver + string(tx.Amount)
	hash := sha256.Sum256([]byte(data))
	r, s, _ := ecdsa.Sign(rand.Reader, privKey, hash[:])
	signature := append(r.Bytes(), s.Bytes()...)
	tx.Signature = hex.EncodeToString(signature)
}

// VerifyTransaction checks if signature is valid
func (tx *Transaction) VerifyTransaction(pubKey ecdsa.PublicKey) bool {
	data := tx.Sender + tx.Receiver + string(tx.Amount)
	hash := sha256.Sum256([]byte(data))

	sigBytes, _ := hex.DecodeString(tx.Signature)
	r := big.Int{}
	s := big.Int{}
	r.SetBytes(sigBytes[:len(sigBytes)/2])
	s.SetBytes(sigBytes[len(sigBytes)/2:])

	return ecdsa.Verify(&pubKey, hash[:], &r, &s)
}
