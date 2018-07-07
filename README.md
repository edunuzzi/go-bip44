# Go-BIP44
[![Go Report Card](https://goreportcard.com/badge/github.com/Swipecoin/go-bip44)](https://goreportcard.com/report/github.com/Swipecoin/go-bip44)

A Golang implementation of the [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki) for Hierarchical Deterministic (HD) addresses. 
It currently only supports Bitcoin, but we plan to add others in the future.

Released under the terms of the [MIT LICENSE](LICENSE).

## Should I use this in production?
This library is in very early stages. Please be aware that some bugs may exist. 

## Can I trust this code?
> Do not trust. Verify.

Since this lib is in its early days, we recommend every user of this library audit and verify any underlying code for its validity and suitability.
You can do so by using [this tool](https://iancoleman.io/bip39/).

## Installation
```bash 
go get -u github.com/Swipecoin/go-bip44 
```

## Usage

### New 24-word Mnemonic and Seed
```golang
// bitSize must be a multiple of 32
bitSize := 256
mnemonic, _ := bip44.NewMnemonic(bitSize)
seedBytes := m.NewSeed("my password")
```

### Master key From Seed Hex
```golang
xKey, _ := bip44.NewKeyFromSeedHex("your secret seed in hex format", bip44.MAINNET)
```

### Master key From Seed bytes
```golang
xKey, _ := bip44.NewKeyFromSeedBytes(seedBytes, bip44.MAINNET)
```

### From base58-encoded Extended Key
```golang
ak, _ := bip44.NewAccountKeyFromXKey(xPubKey)

externalAddress, _ := accountKey.DeriveP2PKAddress(bip44.ExternalChangeType, 0, bip44.MAINNET)
internalAddress, _ := accountKey.DeriveP2PKAddress(bip44.InternalChangeType, 0, bip44.MAINNET)
```

### Deriving P2PK addresses

| coin    | account | chain    | address | path                      |
| ------- | ------- | -------- | ------- | ------------------------- |
| Bitcoin | first   | external | first   | m / 44' / 0' / 0' / 0 / 0 |

```golang 
xKey, _ := bip44.NewKeyFromSeedHex("your secret seed in hex format", bip44.MAINNET)
accountKey, _ := xKey.BIP44AccountKey(bip44.BitcoinCoinType, 0, true)

externalAddress, _ := accountKey.DeriveP2PKAddress(bip44.ExternalChangeType, 0, bip44.MAINNET)
```

---

| coin    | account | chain    | address | path                      |
| ------- | ------- | -------- | ------- | ------------------------- |
| Bitcoin | first   | external | second  | m / 44' / 0' / 0' / 0 / 1 |

```golang 
xKey, _ := bip44.NewKeyFromSeedHex("your secret seed in hex format", bip44.MAINNET)
accountKey, _ := xKey.BIP44AccountKey(bip44.BitcoinCoinType, 0, true)

externalAddress, _ := accountKey.DeriveP2PKAddress(bip44.ExternalChangeType, 1, bip44.MAINNET)
```

---

| coin            | account  | chain    | address | path                      |
| --------------- | -------- | -------- | ------- | ------------------------- |
| Bitcoin Testnet | second   | internal | first   | m / 44' / 1' / 1' / 1 / 0 |

```golang 
xKey, _ := bip44.NewKeyFromSeedHex("your secret seed in hex format", bip44.TESTNET3)
accountKey, _ := xKey.BIP44AccountKey(bip44.TestnetCoinType, 1, true)

externalAddress, _ := accountKey.DeriveP2PKAddress(bip44.InternalChangeType, 0, bip44.TESTNET3)
```

## TODO
- [X] Report badge
- [ ] Unit Tests
- [ ] Create GoDoc
- [ ] Stellar
- [ ] Ethereum

## Contribution
Please feel free to contribute with both suggestions and pull requests :D
