package main

import (
	"fmt"
	"log"
	//"github.com/Roshan310/DaanVeer/blockchain"
	"github.com/Roshan310/DaanVeer/wallet"

)

//--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
//Test for Blockchain And addition of transaction

// func main() {
// 	blockchain := blockchain.NewBlockchain()
// 	blockchain.Print()

// 	blockchain.AddTransaction("A", "B", 1.0)
// 	blockchain.AddTransaction("C", "D", 2.0)
// 	previousHash := blockchain.LastBlock().Hash()
// 	blockchain.CreateBlock(previousHash)
// 	blockchain.Print()

// 	blockchain.AddTransaction("C", "D", 2.0)
// 	previousHash = blockchain.LastBlock().Hash()
// 	blockchain.CreateBlock(previousHash)
// 	blockchain.Print()
// }


//--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Test for PoA
// func main() {
// 	blockchain1 := blockchain.NewBlockchain()
// 	poa := blockchain.NewPoA([]string{"authority1", "authority2", "authority3"})

// 	poa.AddAuthority("authority3")
// 	blockchain1.AddTransaction("A", "B", 1.0)

// 	prevHash := blockchain1.LastBlock().Hash()
// 	block := blockchain.NewBlock(prevHash, blockchain1.TransactionPool)

// 	poa.SignBlock("authority1", block)
// 	blockchain1.Chain = append(blockchain1.Chain, block)
// 	blockchain1.TransactionPool = []*blockchain.Transactions{}

// 	if poa.VerifyBlock(block, "authority1") {
// 		fmt.Println("Block successfully verified by authority1")
// 	} else {
// 		fmt.Println("Block verification failed")
// 	}

// 	blockchain1.Print()
// }
//--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

// Test for Wallet
	func main() {
		// Specify the filename to save the wallet.
		walletFile := "my_wallet.txt"
	
		// Generate a new wallet and save it to the file.
		myWallet, err := wallet.GenerateWallet(walletFile)
		if err != nil {
			log.Fatalf("Failed to generate wallet: %v\n", err)
		}
	
		// Display the newly generated wallet information.
		fmt.Println("New Wallet:")
		fmt.Println("Wallet Address:", myWallet.Address)
		fmt.Printf("Public Key: %x\n", wallet.PublicKeyToBytes(myWallet.PublicKey))
		fmt.Printf("Private Key: %x\n", myWallet.PrivateKey.D.Bytes())
	
		// Load all wallets from the file.
		wallets, err := wallet.LoadAllWallets(walletFile)
		if err != nil {
			log.Fatalf("Failed to load wallets: %v\n", err)
		}
	
		// Display all loaded wallets.
		fmt.Println("\nLoaded Wallets:")
		for i, w := range wallets {
			fmt.Printf("Wallet %d:\n", i+1)
			fmt.Println("Wallet Address:", w.Address)
			fmt.Printf("Public Key: %x\n", wallet.PublicKeyToBytes(w.PublicKey))
			fmt.Printf("Private Key: %x\n\n", w.PrivateKey.D.Bytes())
		}
	}
	



