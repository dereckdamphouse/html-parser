package req

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const htmlStub = "{\"html\":\"<html></html>\",\"properties\":[{\"name\":\"image\",\"selector\":\".image img\",\"attribute\":\"src\"}]}"

func TestUnmarshal(t *testing.T) {
	tt := []struct {
		name        string
		body        string
		unmarshaler func(data []byte, v any) error
		data        *Data
		err         error
	}{
		{
			"handles Unmarshaler error",
			"",
			func(data []byte, v any) error {
				return fmt.Errorf("some error")
			},
			&Data{},
			fmt.Errorf("some error"),
		},
		{
			"handles successful Unmarshaler",
			htmlStub,
			json.Unmarshal,
			&Data{
				HTML: "<html></html>",
				Properties: []Property{{
					Name:      "image",
					Selector:  ".image img",
					Attribute: "src",
				}},
			},
			nil,
		},
		{
			"handles missing html field",
			"{\"properties\":[{\"name\":\"title\",\"selector\":\".title\"}]}",
			json.Unmarshal,
			&Data{
				HTML:       "",
				Properties: []Property(nil),
			},
			fmt.Errorf("request body is missing 'html' field"),
		},
		{
			"handles missing properties field",
			"{\"html\":\"<html></html>\"}",
			json.Unmarshal,
			&Data{
				HTML:       "",
				Properties: []Property(nil),
			},
			fmt.Errorf("request body is missing 'properties' field"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			unmarshaler = tc.unmarshaler
			data, err := Unmarshal(tc.body)
			assert.Equal(t, tc.data, data)
			assert.Equal(t, tc.err, err)
		})
	}
}
