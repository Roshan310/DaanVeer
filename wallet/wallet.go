package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"math/big"
	"strings"
	"encoding/hex"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
)

const (
	CHECK_SUM_LENGTH = 4
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey *ecdsa.PublicKey
	Address string
}

func (w *Wallet) GenerateKeyPair() error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	w.PrivateKey = privateKey
	w.PublicKey = &privateKey.PublicKey
	return nil
}

func PublicKeyToBytes(publicKey *ecdsa.PublicKey) []byte {
	return append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)
}

func BytesToPublicKey(pubKeyBytes []byte) (*ecdsa.PublicKey, error) {
	curve := elliptic.P256()
	keyLen := len(pubKeyBytes)/2
	if keyLen == 0 {
		return nil, errors.New("invalid bytes of public key")
	}
	x := new(big.Int).SetBytes(pubKeyBytes[:keyLen])
	y := new(big.Int).SetBytes(pubKeyBytes[keyLen:])

	if !curve.IsOnCurve(x, y) {
		return nil, errors.New("invalid points on the curve for this public key")
	}

	return &ecdsa.PublicKey{Curve: curve, X: x, Y: y}, nil
}


func PublicKeyHashRipeMD160(pubKey *ecdsa.PublicKey) []byte {

	//Using SHA256 for hashing right here during RipeMD160 hash.
	pubKeybytes := PublicKeyToBytes((pubKey))
	pubKeyHash := sha256.Sum256(pubKeybytes)

	//ripemd160 maintains internal state for the hash so, creating a new object.
	ripeMDHasher := ripemd160.New()
	_, _ = ripeMDHasher.Write(pubKeyHash[:])
	return ripeMDHasher.Sum(nil)
}

func GenerateAddress(publicKey *ecdsa.PublicKey) string {

	//	The address generation process simply follows the following steps:
	//		1. Hash the public key using sha256.
	//		2. Use RipeMD160 hash to convert it to 160 bits.
	//		3. Add checkSum error code
	//		4. Append the public key hash and checksum code
	//		5. Encode using base 58 to get the address

	publicKeyHash := PublicKeyHashRipeMD160(publicKey)
	 
	//here we can add version number as like in bitcoin address but we are not doing that for now
	checkSum := calculateCheckSum(publicKeyHash)
	finalHash := append(publicKeyHash, checkSum...)
	address := base58.Encode(finalHash)

	return address

}

func calculateCheckSum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:CHECK_SUM_LENGTH]
}


// func (w *Wallet) SaveToFile(fileName string) error {
// 	data := fmt.Sprintf("%x\n%s\n%s", w.PrivateKey.D.Bytes(), PublicKeyToBytes(w.PublicKey), w.Address)

// 	return os.WriteFile(fileName, []byte(data), 0600)
// }

// func (w *Wallet) LoadFromFile(fileName string) error {
// 	data, err := os.ReadFile(fileName)

// 	if err != nil {
// 		return err
// 	}

// 	var privKeyBytes []byte
// 	var pubKeyBytes []byte
// 	fmt.Sscanf(string(data), "%x\n%s\n%s", &privKeyBytes, &pubKeyBytes, &w.Address)

// 	//Here, we reconstruct the private key from its bytes
// 	w.PrivateKey = new(ecdsa.PrivateKey)
// 	w.PrivateKey.PublicKey.Curve = elliptic.P256()
// 	w.PrivateKey.D = new(big.Int).SetBytes(privKeyBytes)
// 	w.PrivateKey.PublicKey.X, w.PrivateKey.PublicKey.Y = elliptic.P256().ScalarBaseMult(privKeyBytes)

// 	//Now, reconstructing the public key
// 	w.PublicKey, err = BytesToPublicKey(pubKeyBytes)
// 	if err != nil {
// 		return err
// 	}	

// 	return nil
// }
// func (w *Wallet) SaveToFile(fileName string) error {
// 	// Serialize private key, public key, and address
// 	data := fmt.Sprintf(
// 		"%x\n%x\n%s",
// 		w.PrivateKey.D.Bytes(),
// 		PublicKeyToBytes(w.PublicKey),
// 		w.Address,
// 	)

// 	// Write to file
// 	return os.WriteFile(fileName, []byte(data), 0600)
// }

// func (w *Wallet) LoadFromFile(fileName string) error {
// 	// Read file data
// 	data, err := os.ReadFile(fileName)
// 	if err != nil {
// 		return err
// 	}

// 	// Split file content into lines
// 	lines := strings.Split(string(data), "\n")
// 	if len(lines) < 3 {
// 		return errors.New("invalid wallet file format")
// 	}

// 	// Parse private key
// 	privKeyBytes, err := hex.DecodeString(lines[0])
// 	if err != nil {
// 		return fmt.Errorf("failed to parse private key: %v", err)
// 	}
// 	w.PrivateKey = new(ecdsa.PrivateKey)
// 	w.PrivateKey.PublicKey.Curve = elliptic.P256()
// 	w.PrivateKey.D = new(big.Int).SetBytes(privKeyBytes)
// 	w.PrivateKey.PublicKey.X, w.PrivateKey.PublicKey.Y = elliptic.P256().ScalarBaseMult(privKeyBytes)

// 	// Parse public key
// 	pubKeyBytes, err := hex.DecodeString(lines[1])
// 	if err != nil {
// 		return fmt.Errorf("failed to parse public key: %v", err)
// 	}
// 	w.PublicKey, err = BytesToPublicKey(pubKeyBytes)
// 	if err != nil {
// 		return fmt.Errorf("failed to reconstruct public key: %v", err)
// 	}

// 	// Parse address
// 	w.Address = lines[2]

// 	return nil
// }

func (w *Wallet) SaveToFile(fileName string) error {
	// Serialize private key, public key, and address
	data := fmt.Sprintf(
		"%x\n%x\n%s\n\n",
		w.PrivateKey.D.Bytes(),
		PublicKeyToBytes(w.PublicKey),
		w.Address,
	)

	// Open the file in append mode
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the wallet data to the file
	_, err = file.WriteString(data)
	return err
}

func LoadAllWallets(fileName string) ([]*Wallet, error) {
	// Read file data
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Split the file into individual wallet blocks
	walletBlocks := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	var wallets []*Wallet
	for _, block := range walletBlocks {
		lines := strings.Split(block, "\n")
		if len(lines) < 3 {
			return nil, errors.New("invalid wallet block format")
		}

		// Parse each wallet
		privKeyBytes, err := hex.DecodeString(lines[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %v", err)
		}
		pubKeyBytes, err := hex.DecodeString(lines[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse public key: %v", err)
		}
		address := lines[2]

		// Construct wallet
		wallet := &Wallet{}
		wallet.PrivateKey = new(ecdsa.PrivateKey)
		wallet.PrivateKey.PublicKey.Curve = elliptic.P256()
		wallet.PrivateKey.D = new(big.Int).SetBytes(privKeyBytes)
		wallet.PrivateKey.PublicKey.X, wallet.PrivateKey.PublicKey.Y = elliptic.P256().ScalarBaseMult(privKeyBytes)

		wallet.PublicKey, err = BytesToPublicKey(pubKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to reconstruct public key: %v", err)
		}
		wallet.Address = address

		wallets = append(wallets, wallet)
	}

	return wallets, nil
}



func GenerateWallet(filename string) (*Wallet, error) {
	wallet := &Wallet{}
	if err := wallet.GenerateKeyPair(); err != nil {
		return nil, err
	}
	wallet.Address = GenerateAddress(wallet.PublicKey)
	if err := wallet.SaveToFile(filename); err != nil {
		return nil, err
	}

	fmt.Printf("Your wallet is generated and here is your address %s\n", wallet.Address)
	return wallet, nil
}