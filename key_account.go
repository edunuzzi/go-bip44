package bitcoin_address

import (
	"github.com/btcsuite/btcutil/hdkeychain"
)

type KeyAccount struct {
	*hdkeychain.ExtendedKey
	HDStartPath
}

func (k *KeyAccount) DeriveAddress(changeType ChangeType, index uint32, network Network) (*Address, error) {

	var changeTypeIndex = uint32(changeType)

	if k.ExtendedKey.IsPrivate() {
		changeType += HardenedKeyZeroIndex
	}
	if k.ExtendedKey.IsPrivate() {
		index += HardenedKeyZeroIndex
	}

	changeTypeK, err := k.ExtendedKey.Child(changeTypeIndex)
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

		},
		Value: a.EncodeAddress(),
	}, nil
}
