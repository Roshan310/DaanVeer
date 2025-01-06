package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
	"crypto/sha256"
)


type Transactions struct {
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transactions {
	return &Transactions{SenderBlockchainAddress: sender, RecipientBlockchainAddress: recipient, Value: value}
}

func (t *Transactions) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("Sender Blockchain Address:    %s\n", t.SenderBlockchainAddress)
	fmt.Printf("Recipient Blockchain Address: %s\n", t.RecipientBlockchainAddress)
	fmt.Printf("Value:                        %.1f\n", t.Value)
}

func (t *Transactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchainaddress"`
		Recipient string  `json:"recipient_blockchainaddress"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.SenderBlockchainAddress,
		Recipient: t.RecipientBlockchainAddress,
		Value:     t.Value,
	})
}

func (t *Transactions) Hash() [32]byte {
	m, _ := json.Marshal(t)
	return sha256.Sum256(m)
}