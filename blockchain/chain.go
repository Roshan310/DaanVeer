package blockchain

import(
	"fmt"
	"strings"
)
type Blockchain struct {
	TransactionPool []*Transactions
	Chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(b.Hash()) // Genesis block
	return bc
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.TransactionPool = append(bc.TransactionPool, t)
}

func (bc *Blockchain) Print() {
	for i, block := range bc.Chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) CreateBlock(previousHash [32]byte) *Block {
	b := NewBlock(previousHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, b)
	bc.TransactionPool = []*Transactions{}
	return b
}