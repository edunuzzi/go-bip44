package main

import "github.com/Swipecoin/bitcoin_address"

func main() {
	xKey, err := bitcoin_address.NewKeyFromSeedHex("", bitcoin_address.MAINNET)

	accountKey, err := xKey.BIP44AccountKey(bitcoin_address.BitcoinCoinType, 0, true)

	internalAddress, err := accountKey.DeriveAddress(bitcoin_address.InternalChangeType, 0, bitcoin_address.MAINNET)
	internalAddress, err := accountKey.DeriveAddress(bitcoin_address.ExternalChangeType, 0, bitcoin_address.MAINNET)

	accountKey, err = xKey.BIP49AccountKey(bitcoin_address.BitcoinCoinType, 0, true)

	internalAddress, err := accountKey.DeriveAddress(bitcoin_address.InternalChangeType, 0, bitcoin_address.TESTNET3)
	internalAddress, err := accountKey.DeriveAddress(bitcoin_address.ExternalChangeType, 0, bitcoin_address.TESTNET3)

	accountKey, err = xKey.BIP84AccountKey(bitcoin_address.BitcoinCoinType, 0, true)

	internalAddress, err := accountKey.DeriveAddress(bitcoin_address.InternalChangeType, 0, bitcoin_address.TESTNET3)
	internalAddress, err := accountKey.DeriveAddress(bitcoin_address.ExternalChangeType, 0, bitcoin_address.TESTNET3)
}
