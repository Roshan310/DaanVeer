package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strings"
	"math/big"

	"github.com/Roshan310/DaanVeer/wallet"
)


type Transactions struct {
	SenderHash    []byte
	RecipientHash []byte
	Value         float32
	Signature     []byte
	Timestamp     uint64
}

func NewTransaction(sender []byte, recipient []byte, value float32) *Transactions {
	return &Transactions{SenderHash: sender, RecipientHash: recipient, Value: value}
}

func (t *Transactions) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("Sender Address:    %s\n", t.SenderHash)
	fmt.Printf("Recipient Address: %s\n", t.RecipientHash)
	fmt.Printf("Value:                        %.1f\n", t.Value)
	fmt.Printf("Signature: %s", t.Signature)
	fmt.Printf("Timestamp %d", t.Timestamp)
}

func (t *Transactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_address"`
		Recipient string  `json:"recipient_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    string(t.SenderHash),
		Recipient: string(t.RecipientHash),
		Value:     t.Value,
	})
}

func (t *Transactions) Hash() []byte {
	m, _ := json.Marshal(t)
	hash := sha256.Sum256(m)
	return hash[:]
}

func (tx Transaction) SerializeTransaction() ([]byte, error) {
	var encoded bytes.Buffer
	err := gob.NewEncoder(&encoded).Encode(tx)
	return encoded.Bytes(), err
}

func DeserializeTransaction(serializedTransaction []byte) (*Transaction, error) {
	var transaction Transaction
	err := gob.NewDecoder(bytes.NewReader(serializedTransaction)).Decode((&transaction))
	return &transaction, err
}

// func CoinBaseTransaction(srcWallet *wallet.Wallet, amount uint64, chain *Blockchain) (*Transactions, error) {
// 	// check if the address is valid
// 	walletAddress := string(srcWallet.Address)
// 	pubKeyHash, err := wallet.PubKeyFromAddress(walletAddress)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

func (tx *Transactions) SignTransaction(wallet *wallet.Wallet) error {
	senderPrivKey := wallet.PrivateKey

	r, s, err := ecdsa.Sign(rand.Reader, senderPrivKey, tx.Hash())
	if err != nil {
		fmt.Println("failed to sign the transaction!!!")
	}
	signature := append(r.Bytes(), s.Bytes()...)
	tx.Signature = signature
	return nil
}
func (tx *Transactions) VerifyTransaction(pubKey *ecdsa.PublicKey) bool {
	signatureBytes := tx.Signature

	r := new(big.Int).SetBytes(signatureBytes[:len(signatureBytes)/2])
	s := new(big.Int).SetBytes(signatureBytes[len(signatureBytes)/2:])

	hash := tx.Hash()

	isValid := ecdsa.Verify(pubKey, hash[:], r, s)
	return isValid
}