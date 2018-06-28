# BitcoinHDAddress
[![Go Report Card](https://goreportcard.com/badge/github.com/algoGuy/EasyMIDI)](https://goreportcard.com/report/github.com/algoGuy/EasyMIDI)
[![GoDoc](https://godoc.org/github.com/algoGuy/EasyMIDI?status.svg)](https://godoc.org/github.com/algoGuy/EasyMIDI)

A Golang implementation of the BIP44, BIP49 and BIP84 for Hierarchical Deterministic (HD) Bitcoin addresses.

Released under the terms of the [MIT LICENSE](LICENSE).  

## Should i use this in production?
This library is in very early stages. Please be aware that some bugs may exist. 

## Can I trust this code?
> Do not trust. Please verify.

We recommend every user of this library audit and verify any underlying code for its validity and suitability.

## Installation
```bash 
go get -u ####TODO 
```

## Quick Start:

### BIP44
m / purpose' / coin_type' / account' / change / address_index

| coin    | account | chain    | address | path                      |
| ------- | ------- | -------- | ------- | ------------------------- |
| Bitcoin | first   | external | first   | m / 44' / 0' / 1' / 1 / 0 |

```golang 
xKey, _ := bitcoin_address.NewKeyFromSeedHex("your secret seed in hex format", bitcoin_address.MAINNET)
accountKey, _ := xKey.BIP44AccountKey(bitcoin_address.BitcoinCoinType, 0, true)

externalAddress, _ := accountKey.DeriveAddress(bitcoin_address.InternalChangeType, 0, bitcoin_address.MAINNET)
```

---

| coin            | account  | chain    | address | path                      |
| --------------- | -------- | -------- | ------- | ------------------------- |
| Bitcoin Testnet | second   | external | first   | m / 44' / 1' / 1' / 0 / 0 |

```golang 
xKey, _ := bitcoin_address.NewKeyFromSeedHex("your secret seed in hex format", bitcoin_address.MAINNET)
accountKey, _ := xKey.BIP44AccountKey(bitcoin_address.TestnetCoinType, 1, true)

externalAddress, _ := accountKey.DeriveAddress(bitcoin_address.ExternalChangeType, 0, bitcoin_address.MAINNET)
```

### BIP49

| coin    | account | chain    | address | path                      |
| ------- | ------- | -------- | ------- | ------------------------- |
| Bitcoin | first   | external | first   | m / 49' / 0' / 1' / 1 / 0 |

```golang 
xKey, _ := bitcoin_address.NewKeyFromSeedHex("your secret seed in hex format", bitcoin_address.MAINNET)
accountKey, _ := xKey.BIP49AccountKey(bitcoin_address.BitcoinCoinType, 0, true)

externalAddress, _ := accountKey.DeriveAddress(bitcoin_address.InternalChangeType, 0, bitcoin_address.MAINNET)
```

---

| coin            | account  | chain    | address | path                      |
| --------------- | -------- | -------- | ------- | ------------------------- |
| Bitcoin Testnet | second   | external | first   | m / 49' / 1' / 1' / 0 / 0 |

```golang 
xKey, _ := bitcoin_address.NewKeyFromSeedHex("your secret seed in hex format", bitcoin_address.MAINNET)
accountKey, _ := xKey.BIP49AccountKey(bitcoin_address.TestnetCoinType, 1, true)

externalAddress, _ := accountKey.DeriveAddress(bitcoin_address.ExternalChangeType, 0, bitcoin_address.MAINNET)
```

### BIP84

| coin    | account | chain    | address | path                      |
| ------- | ------- | -------- | ------- | ------------------------- |
| Bitcoin | first   | external | first   | m / 84' / 0' / 1' / 1 / 0 |
```golang 
xKey, _ := bitcoin_address.NewKeyFromSeedHex("your secret seed in hex format", bitcoin_address.MAINNET)
accountKey, _ := xKey.BIP84AccountKey(bitcoin_address.BitcoinCoinType, 0, true)

externalAddress, _ := accountKey.DeriveAddress(bitcoin_address.InternalChangeType, 0, bitcoin_address.MAINNET)
```

---

| coin            | account  | chain    | address | path                      |
| --------------- | -------- | -------- | ------- | ------------------------- |
| Bitcoin Testnet | second   | external | first   | m / 84' / 1' / 1' / 0 / 0 |
```golang 
xKey, _ := bitcoin_address.NewKeyFromSeedHex("your secret seed in hex format", bitcoin_address.MAINNET)
accountKey, _ := xKey.BIP84AccountKey(bitcoin_address.TestnetCoinType, 1, true)

externalAddress, _ := accountKey.DeriveAddress(bitcoin_address.ExternalChangeType, 0, bitcoin_address.MAINNET)
```


## TODO
- [ ] Unit Tests
- [ ] Create GoDoc
