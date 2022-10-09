package req

import (
	"encoding/json"
	"fmt"
)

var unmarshaler = json.Unmarshal

func Unmarshal(body string) (*Data, error) {
	var d Data

	if err := unmarshaler([]byte(body), &d); err != nil {
		return &Data{}, err
	}

	if d.HTML == "" {
		return &Data{}, fmt.Errorf("request body is missing 'html' field")
	}

	if len(d.Properties) == 0 {
		return &Data{}, fmt.Errorf("request body is missing 'properties' field")
	}

	return &d, nil
}
