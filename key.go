package bitcoin_address

import (
	"github.com/btcsuite/btcutil/hdkeychain"
	"encoding/hex"
)

type ExtendedKey struct {
	*hdkeychain.ExtendedKey
}

func NewKeyFromSeedHex(seed string, net Network) (privateKey *hdkeychain.ExtendedKey, err error) {

	pk, err := hex.DecodeString(seed)

	if err != nil {
		return nil, err
	}

	return NewKeyFromSeedBytes(pk, net)
}

func NewKeyFromSeedBytes(seed []byte, net Network) (privateKey *hdkeychain.ExtendedKey, err error) {

	n, err := networkToChainConfig(net)
	return hdkeychain.NewMaster(seed, n)
}

func NewKeyFromString(value string) (*ExtendedKey, error) {

	xKey, err := hdkeychain.NewKeyFromString(value)

	if err != nil {
		return nil, err
	}

	return &ExtendedKey{xKey}, nil
}

func (e *ExtendedKey) BIP44AccountKey(accIndex uint32, coinType CoinType, includePrivateKey bool) (*KeyAccount, error) {

	return e.baseDeriveAccount(accIndex, BIP44Purpose, coinType, includePrivateKey)
}

func (e *ExtendedKey) BIP49AccountKey(accIndex uint32, coinType CoinType, includePrivateKey bool) (*KeyAccount, error) {

	return e.baseDeriveAccount(accIndex, BIP49Purpose, coinType, includePrivateKey)
}

func (e *ExtendedKey) BIP84AccountKey(accIndex uint32, coinType CoinType, includePrivateKey bool) (*KeyAccount, error) {

	return e.baseDeriveAccount(accIndex, BIP84Purpose, coinType, includePrivateKey)
}

func (e *ExtendedKey) baseDeriveAccount(accIndex uint32, purpose Purpose, coinType CoinType, includePrivateKey bool) (*KeyAccount, error) {

	var purposeIndex = uint32(purpose)
	var coinTypeIndex = uint32(coinType)

	if e.ExtendedKey.IsPrivate() { purposeIndex += HardenedKeyZeroIndex }
	if e.ExtendedKey.IsPrivate() { coinTypeIndex += HardenedKeyZeroIndex }
	if e.ExtendedKey.IsPrivate() { accIndex += HardenedKeyZeroIndex }

	purposeK, err := e.ExtendedKey.Child(purposeIndex)
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
		PurposeIndex: purposeIndex,
		CoinTypeIndex: coinTypeIndex,
		AccountIndex: accIndex,
	}

	if includePrivateKey {
		return &KeyAccount{
			ExtendedKey: accK,
			HDStartPath: hdStartPath,
		}, nil
	}

	pub, err := accK.Neuter()
	if err != nil {
		return nil, err
	}

	return &KeyAccount{
		ExtendedKey: pub,
		HDStartPath: hdStartPath,
	}, nil
}