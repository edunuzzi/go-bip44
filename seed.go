package bip44

import (
	"github.com/tyler-smith/go-bip39"
)

type Mnemonic struct {
	Value   string
}

// bitSize must be a multiple of 32
func NewMnemonic(bitSize int) (*Mnemonic, error) {
	entropy, e := bip39.NewEntropy(bitSize)

	if e != nil {
		return nil, e
	}

	m, e := bip39.NewMnemonic(entropy)

	return &Mnemonic{m}, e
}

func ParseMnemonic(mnemonic string) Mnemonic {
	return Mnemonic{mnemonic}
}

func (m Mnemonic) NewSeed(password string) ([]byte, error) {
	return bip39.NewSeedWithErrorChecking(m.Value, password)
}
