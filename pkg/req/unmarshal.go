package req

import (
	"encoding/json"
	"fmt"
)

type Selector struct {
	Key       string
	Selector  string
	Attribute string
}

type Body struct {
	HTML      string
	Selectors []Selector
}

var unmarshaler = json.Unmarshal

func Unmarshal(reqBody string) (*Body, error) {
	var b Body

	if err := unmarshaler([]byte(reqBody), &b); err != nil {
		return &b, err
	}

	if b.HTML == "" {
		return &b, fmt.Errorf("request body is missing 'html' field")
	}

	if len(b.Selectors) == 0 {
		return &b, fmt.Errorf("request body is missing 'selectors' field")
	}

	return &b, nil
}
