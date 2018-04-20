package bitcoin_address

import (
	"github.com/btcsuite/btcutil/hdkeychain"
)

type KeyAccount struct {
	extendedKey *hdkeychain.ExtendedKey
	HDStartPath
}

func NewKeyAccountFromXKey(value string) (*KeyAccount, error) {
	xKey, err := hdkeychain.NewKeyFromString(value)

	if err != nil {
		return nil, err
	}

	return &KeyAccount{
		extendedKey: xKey,
		HDStartPath: HDStartPath{
			PurposeIndex: -1,
			CoinTypeIndex: -1,
			AccountIndex: -1,
		},
	}, nil
}

func (k *KeyAccount) DeriveAddress(changeType ChangeType, index uint32, network Network) (*Address, error) {

	var changeTypeIndex = uint32(changeType)

	if k.extendedKey.IsPrivate() { changeType = HardenedKeyZeroIndex + changeType }
	if k.extendedKey.IsPrivate() { index = HardenedKeyZeroIndex + index }

	changeTypeK, err := k.extendedKey.Child(changeTypeIndex)
	if err != nil {
		return nil, err
	}

	addressK, err := changeTypeK.Child(index)
	if err != nil {
		return nil, err
	}

	netParam, err := networkToChainConfig(network)

	if err != nil {
		return nil, err
	}

	a, err := addressK.Address(netParam)

	if err != nil {
		return nil, err
	}

	return &Address{
		HDStartPath: HDStartPath{
			PurposeIndex: k.PurposeIndex,
			CoinTypeIndex: k.CoinTypeIndex,
			AccountIndex: k.AccountIndex,
		},
		HDEndPath: HDEndPath{
			ChangeIndex: changeTypeIndex,
			AddressIndex: index,
		},
		Value: a.EncodeAddress(),
	}, nil
}