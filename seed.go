package bitcoin_address

import "github.com/tyler-smith/go-bip39"

type Mnemonic struct {
	Value string
	BitSize int
}

// bitSize must be a multiple of 32
func NewMnemonic(bitSize int) (*Mnemonic, error) {
	entropy, e := bip39.NewEntropy(bitSize)

	if e != nil {
		return nil, e
	}

	m, e := bip39.NewMnemonic(entropy)

	return &Mnemonic{m, bitSize}, e
}

func (m *Mnemonic) NewSeed(password string) []byte {
	return bip39.NewSeed(m.Value, password)
}