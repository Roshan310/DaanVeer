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
	PreviousHash []byte
	Timestamp    int64
	Transactions []Transactions
	MerkleRoot   []byte
	Signature	string
	TxMerkleTree *MerkleTree
}

func NewBlock(previousHash []byte, transactions []Transactions) *Block {
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

func (b *Block) Hash() []byte {
	m, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error while marshalling!!!")
	}
	fmt.Println(string(m))
	hash := sha256.Sum256(m)
	return hash[:]
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		PreviousHash []byte       `json:"previous_hash"`
		MerkleRoot   []byte       `json:"merkle_root"`
		Transactions []Transactions `json:"transactions"`
	}{
		Timestamp:    b.Timestamp,
		PreviousHash: b.PreviousHash,
		MerkleRoot:   b.MerkleRoot,
		Transactions: b.Transactions,
	})
}

func (b *Block) AddTxToBlock(txPool []Transactions) error {
	var tree *MerkleTree
	tree = NewMerkleTree(txPool)
	b.TxMerkleTree = tree 
	return nil
}
