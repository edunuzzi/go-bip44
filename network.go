package bip44

import (
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
)

type Network int16

const (
	TESTNET3 Network = 0
	MAINNET  Network = 1
)

func networkToChainConfig(net Network) (*chaincfg.Params, error) {
	switch net {
	case TESTNET3:
		return &chaincfg.TestNet3Params, nil

	case MAINNET:
		return &chaincfg.MainNetParams, nil
	}

	return nil, errors.New("invalid network")
}
