package bip44

import (
	"fmt"
	"strconv"
)

type URIParams struct {
	Address string
	Amount  float64
	Label   string
	Message string
}

func EncodeURI(p URIParams) (string, error) {

	// TODO check if valid address
	if p.Address == "" {
		return "", fmt.Errorf("invalid address")
	}

	if p.Amount == 0 {
		return "", fmt.Errorf("invalid amount '0'")
	}

	var uri = fmt.Sprintf(
		"bitcoin:%s?amount=%s",
		p.Address,
		strconv.FormatFloat(p.Amount, 'f', -1, 64),
	)

	if p.Label != "" {
		uri = uri + fmt.Sprintf("&label=%s", p.Label)
	}

	if p.Message != "" {
		uri = uri + fmt.Sprintf("&message=%s", p.Message)
	}

	return uri, nil
}
