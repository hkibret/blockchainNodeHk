package main

import "fmt"

func main() {
	bc := Blockchain{}
	bc.Blocks = append(bc.Blocks, CreateGenesisBlock())

	// Generate keys for Alice
	alicePriv, _ := GenerateKeyPair()

	// Create transactions
	tx1 := Transaction{Sender: "Alice", Receiver: "Bob", Amount: 10}
	tx1.SignTransaction(alicePriv)

	tx2 := Transaction{Sender: "Alice", Receiver: "Charlie", Amount: 5}
	tx2.SignTransaction(alicePriv)

	// Add transactions to mempool
	bc.AddTransaction(tx1)
	bc.AddTransaction(tx2)

	fmt.Println("Mempool size:", len(bc.Mempool))

	// Mine block with pending transactions
	bc.MineBlock()

	// Print blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\nHash: %s\nPrevHash: %s\nTransactions: %+v\n\n",
			block.Index, block.Hash, block.PrevHash, block.Transactions)
	}

	fmt.Println("Blockchain valid?", bc.IsValid())
}
