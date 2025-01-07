package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"errors"
	"math/big"
)

type wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey *ecdsa.PublicKey
	Address string
}

func (w *wallet) GenerateKeyPair() error {
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
	pubKeybytes := PublicKeyToBytes((pubKey))
	pubKeyHash := sha256.Sum256(pubKeybytes)

	ripeMDHasher := ripemd160.New()
	_, _ = ripeMDHasher.Write(pubKeyHash[:])
	return ripeMDHasher.Sum(nil)
}

// func GenerateAddress(publicKey *ecdsa.PublicKey) string {
// 	publicKeyHash := PublicKeyHashRipeMD160(publicKey)
	 
// 	//here we can add version number as like in bitcoin address but we are not doing that for now

// 	checkSum := calculateCheckSum(publicKeyHash)
// 	finalHash := append(publicKeyHash, checkSum...)
// 	return 
// }