package bip44

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/hdkeychain"
)

type ExtendedKey struct {
	key *hdkeychain.ExtendedKey
}

func NewKeyFromSeedHex(seed string, net Network) (*ExtendedKey, error) {

	pk, err := hex.DecodeString(seed)

	if err != nil {
		return nil, err
	}

	return NewKeyFromSeedBytes(pk, net)
}

func NewKeyFromSeedBytes(seed []byte, net Network) (*ExtendedKey, error) {

	n, err := networkToChainConfig(net)

	if err != nil {
		return nil, err
	}

	xKey, err := hdkeychain.NewMaster(seed, n)

	if err != nil {
		return nil, err
	}

	return &ExtendedKey{
		key: xKey,
	}, nil
}

func (e *ExtendedKey) BIP44AccountKey(coinType CoinType, accIndex uint32, includePrivateKey bool) (*AccountKey, error) {

	return e.baseDeriveAccount(BIP44Purpose, coinType, accIndex, includePrivateKey)
}

func (e *ExtendedKey) baseDeriveAccount(purpose Purpose, coinType CoinType, accIndex uint32, includePrivateKey bool) (*AccountKey, error) {

	var purposeIndex = uint32(purpose)
	var coinTypeIndex = uint32(coinType)

	if e.key.IsPrivate() {
		purposeIndex = HardenedKeyZeroIndex + purposeIndex
		coinTypeIndex = HardenedKeyZeroIndex + coinTypeIndex
		accIndex = HardenedKeyZeroIndex + accIndex
	}

	purposeK, err := e.key.Child(purposeIndex)
	if err != nil {
		return nil, err
	}

	cTypeK, err := purposeK.Child(coinTypeIndex)
	if err != nil {
		return nil, err
	}

	accK, err := cTypeK.Child(accIndex)
	if err != nil {
		return nil, err
	}

	hdStartPath := HDStartPath{
		PurposeIndex:  purposeIndex,
		CoinTypeIndex: coinTypeIndex,
		AccountIndex:  accIndex,
	}

	if includePrivateKey {
		return &AccountKey{
			extendedKey: accK,
			startPath:   hdStartPath,
		}, nil
	}

	pub, err := accK.Neuter()
	if err != nil {
		return nil, err
	}

	return &AccountKey{
		extendedKey: pub,
		startPath:   hdStartPath,
	}, nil
}
