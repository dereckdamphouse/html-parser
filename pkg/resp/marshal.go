package resp

import (
	"encoding/json"

	"github.com/dereckdamphouse/html-parser/pkg/html"
)

var marshaler = json.Marshal

func Marshal(parsed html.Parsed) (string, error) {
	res, err := marshaler(parsed)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
