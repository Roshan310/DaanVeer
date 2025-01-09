package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct {
	PreviousHash [32]byte
	Timestamp    int64
	Transactions []*Transactions
	MerkleRoot   [32]byte
	Signature	string
}

func NewBlock(previousHash [32]byte, transactions []*Transactions) *Block {
	b := new(Block)
	b.Timestamp = time.Now().UnixNano()
	b.PreviousHash = previousHash
	b.Transactions = transactions
	merkleTree := NewMerkleTree(transactions)
	b.MerkleRoot = merkleTree.CalculateMerkleRoot()
	return b
}


func (b *Block) Print() {
	fmt.Printf("Timestamp:       %d\n", b.Timestamp)
	fmt.Printf("Previous Hash:   %x\n", b.PreviousHash)
	fmt.Printf("Merkle Root:     %x\n", b.MerkleRoot)
	for _, t := range b.Transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		PreviousHash [32]byte       `json:"previous_hash"`
		MerkleRoot   [32]byte       `json:"merkle_root"`
		Transactions []*Transactions `json:"transactions"`
	}{
		Timestamp:    b.Timestamp,
		PreviousHash: b.PreviousHash,
		MerkleRoot:   b.MerkleRoot,
		Transactions: b.Transactions,
	})
}

