package main

import(
	"blockchain.com/project/blockchain"
)
func main() {
	blockchain := blockchain.NewBlockchain()
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(previousHash)
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(previousHash)
	blockchain.Print()
}