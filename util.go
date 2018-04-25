package bitcoin_address

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

func EncodeURI(p URIParams) string {
	return fmt.Sprintf(
		"bitcoin:%s?amount=%s&label=%s&message=%s",
		p.Address,
		strconv.FormatFloat(p.Amount, 'f', -1, 64),
		p.Label,
		p.Message,
	)
}
