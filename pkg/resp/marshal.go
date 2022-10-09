package resp

import "encoding/json"

type Body struct {
	Error string `json:"error"`
	Found any    `json:"found"`
}

var marshaler = json.Marshal

func Marshal(found any) (string, error) {
	res, err := marshaler(Body{Found: found})
	if err != nil {
		return "", err
	}

	return string(res), nil
}
