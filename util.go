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

//FIXME return error if not amount or address
func EncodeURI(p URIParams) string {

	var uri = fmt.Sprintf(
		"bitcoin:%s?amount=%s",
		p.Address,
		strconv.FormatFloat(p.Amount, 'f', -1, 64),
	)

	if p.Label != "" {
		uri = uri + fmt.Sprintf("&label=%s", p.Label)
	}

	if p.Message != "" {
		uri = uri + fmt.Sprintf("&message=%s", p.Label)
	}

	return uri
}
