package main

import (
	"fmt"

	"blockchain.com/project/blockchain"
)

//Test for Blockchain And addition of transaction

// func main() {
// 	blockchain := blockchain.NewBlockchain()
// 	blockchain.Print()

// 	blockchain.AddTransaction("A", "B", 1.0)
// 	previousHash := blockchain.LastBlock().Hash()
// 	blockchain.CreateBlock(previousHash)
// 	blockchain.Print()

// 	blockchain.AddTransaction("C", "D", 2.0)
// 	previousHash = blockchain.LastBlock().Hash()
// 	blockchain.CreateBlock(previousHash)
// 	blockchain.Print()
// }
// Test for PoA
func main() {
	blockchain1 := blockchain.NewBlockchain()
	poa := blockchain.NewPoA([]string{"authority1", "authority2", "authority3"})

	poa.AddAuthority("authority4")
	blockchain1.AddTransaction("A", "B", 1.0)

	prevHash := blockchain1.LastBlock().Hash()
	block := blockchain.NewBlock(prevHash, blockchain1.TransactionPool)

	poa.SignBlock("authority1", block)
	blockchain1.Chain = append(blockchain1.Chain, block)
	blockchain1.TransactionPool = []*blockchain.Transactions{}

	if poa.VerifyBlock(block, "authority1") {
		fmt.Println("Block successfully verified by authority1")
	} else {
		fmt.Println("Block verification failed")
	}

	blockchain1.Print()
}



