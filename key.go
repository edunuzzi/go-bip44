package bitcoin_address

import (
	"github.com/btcsuite/btcutil/hdkeychain"
	"encoding/hex"
)

type ExtendedKey struct {
	extendedKey *hdkeychain.ExtendedKey
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

	return e.baseDeriveAccount(BIP44Purpose, accIndex, coinType, includePrivateKey)
}

func (e *ExtendedKey) BIP49AccountKey(accIndex uint32, coinType CoinType, includePrivateKey bool) (*KeyAccount, error) {

	return e.baseDeriveAccount(BIP49Purpose, accIndex, coinType, includePrivateKey)
}

func (e *ExtendedKey) BIP84AccountKey(accIndex uint32, coinType CoinType, includePrivateKey bool) (*KeyAccount, error) {

	return e.baseDeriveAccount(BIP84Purpose, accIndex, coinType, includePrivateKey)
}

func (e *ExtendedKey) baseDeriveAccount(purpose Purpose, accIndex uint32, coinType CoinType, includePrivateKey bool) (*KeyAccount, error) {

	var purposeIndex = uint32(purpose)
	var coinTypeIndex = uint32(coinType)

	if e.extendedKey.IsPrivate() { purposeIndex = HardenedKeyZeroIndex + purposeIndex }
	if e.extendedKey.IsPrivate() { coinTypeIndex = HardenedKeyZeroIndex + coinTypeIndex }
	if e.extendedKey.IsPrivate() { accIndex = HardenedKeyZeroIndex + accIndex }

	purposeK, err := e.extendedKey.Child(purposeIndex)
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
		PurposeIndex: int32(purposeIndex),
		CoinTypeIndex: int32(coinTypeIndex),
		AccountIndex: int32(accIndex),
	}

	if includePrivateKey {
		return &KeyAccount{
			extendedKey: accK,
			HDStartPath: hdStartPath,
		}, nil
	}

	pub, err := accK.Neuter()
	if err != nil {
		return nil, err
	}

	return &KeyAccount{
		extendedKey: pub,
		HDStartPath: hdStartPath,
	}, nil
}