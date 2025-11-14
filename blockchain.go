package main

import (
	"fmt"
	"strings"
	"time"
)

type Blockchain struct {
	Blocks []Block
	// Mempool   []Transaction
	Mempool []Transaction
}

// AddTransaction adds a transaction to the mempool
func (bc *Blockchain) AddTransaction(tx Transaction) {
	bc.Mempool = append(bc.Mempool, tx)
}

// MineBlock creates a new block with all pending transactions
func (bc *Blockchain) MineBlock() {
	if len(bc.Mempool) == 0 {
		fmt.Println("No transactions to mine")
		return
	}

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now(),
		Transactions: bc.Mempool,
		PrevHash:     prevBlock.Hash,
		Nonce:        0,
	}
	newBlock.Hash = newBlock.CalculateHash()
	bc.Blocks = append(bc.Blocks, newBlock)

	// Clear mempool after mining
	bc.Mempool = []Transaction{}
}
func CreateGenesisBlock() Block {
	genesis := Block{Index: 0, Timestamp: time.Now(), Data: "Genesis Block", PrevHash: ""}
	genesis.Hash = genesis.CalculateHash()
	return genesis
}

func (bc *Blockchain) AddBlock(transactions []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
	}
	newBlock.Hash = newBlock.CalculateHash()
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		curr := bc.Blocks[i]
		prev := bc.Blocks[i-1]

		if curr.Hash != curr.CalculateHash() {
			return false
		}
		if curr.PrevHash != prev.Hash {
			return false
		}
	}
	return true
}

const difficulty = 3 // number of leading zeros required

func (b *Block) MineBlock() {
	target := strings.Repeat("0", difficulty)
	for {
		b.Hash = b.CalculateHash()
		if b.Hash[:difficulty] == target {
			fmt.Println("Block mined:", b.Hash)
			break
		}
		b.Nonce++
	}
}
