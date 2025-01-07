package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Authority struct {
	Address string
	IsValid bool
}

type PoA struct {
	Authorities []Authority
}

func NewPoA(addresses []string) *PoA {
	authorities := make([]Authority, len(addresses))
	for i, addr := range addresses {
		authorities[i] = Authority{Address: addr, IsValid: true}
	}
	return &PoA{Authorities: authorities}
}

func (poa *PoA) IsAuthorized(address string) bool {
	for _, authority := range poa.Authorities {
		if authority.Address == address && authority.IsValid {
			return true
		}
	}
	return false
}

func (poa *PoA) AddAuthority(address string) {
	poa.Authorities = append(poa.Authorities, Authority{Address: address, IsValid: true})
}

func (poa *PoA) RevokeAuthority(address string) {
	for i, authority := range poa.Authorities {
		if authority.Address == address {
			poa.Authorities[i].IsValid = false
			break
		}
	}
}

func (poa *PoA) SignBlock(authorityAddress string, block *Block) {
	if !poa.IsAuthorized(authorityAddress) {
		panic("Unauthorized authority attempted to sign block")
	}
	block.Signature = poa.GenerateSignature(authorityAddress, block)
}

func (poa *PoA) GenerateSignature(authorityAddress string, block *Block) string {
	data := fmt.Sprintf("%x|%s|%d", block.PreviousHash, authorityAddress, block.Timestamp)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

func (poa *PoA) VerifyBlock(block *Block, authorityAddress string) bool {
	if !poa.IsAuthorized(authorityAddress) {
		return false
	}
	generatedSignature := poa.GenerateSignature(authorityAddress, block)
	return block.Signature == generatedSignature
}

// Extending the Block struct
// type Block struct {
// 	PreviousHash [32]byte
// 	Timestamp    int64
// 	Transactions []*Transactions
// 	MerkleRoot   [32]byte
// 	Signature    string
// }

// func NewBlock(previousHash [32]byte, transactions []*Transactions) *Block {
// 	b := new(Block)
// 	b.Timestamp = time.Now().UnixNano()
// 	b.PreviousHash = previousHash
// 	b.Transactions = transactions
// 	b.MerkleRoot = CalculateMerkleRoot(transactions)
// 	return b
// }

// Sample usage in main
// func main() {
// 	// Create a PoA system
// 	addresses := []string{"authority1", "authority2", "authority3"}
// 	poa := NewPoA(addresses)

// 	// Create a blockchain
// 	blockchain := NewBlockchain()

// 	// Create and sign a block using an authorized authority
// 	authority := "authority1"
// 	if poa.IsAuthorized(authority) {
// 		block := NewBlock(blockchain.LastBlock().Hash(), []*Transactions{})
// 		poa.SignBlock(authority, block)
// 		blockchain.Chain = append(blockchain.Chain, block)
// 		fmt.Println("Block signed and added to the chain")
// 	} else {
// 		fmt.Println("Authority is not authorized")
// 	}

// 	// Verify the block by another authority
// 	verifyingAuthority := "authority2"
// 	if poa.VerifyBlock(blockchain.LastBlock(), authority) {
// 		fmt.Println("Block verified successfully by", verifyingAuthority)
// 	} else {
// 		fmt.Println("Block verification failed by", verifyingAuthority)
// 	}
//}
