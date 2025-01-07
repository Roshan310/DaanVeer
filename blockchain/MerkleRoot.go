package blockchain
import	(
	"crypto/sha256"
)
func CalculateMerkleRoot(transactions []*Transactions) [32]byte {
	if len(transactions) == 0 {
		return [32]byte{}
	}

	hashes := make([][32]byte, len(transactions))
	for i, t := range transactions {
		hashes[i] = t.Hash()
	}

	for len(hashes) > 1 {
		if len(hashes)%2 != 0 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}

		newLevel := make([][32]byte, 0)
		for i := 0; i < len(hashes); i += 2 {
			h := sha256.Sum256(append(hashes[i][:], hashes[i+1][:]...))
			newLevel = append(newLevel, h)
		}
		hashes = newLevel
	}

	return hashes[0]
}
