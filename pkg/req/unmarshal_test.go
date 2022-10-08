package req

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const htmlStub = "{\"html\":\"<html></html>\",\"selectors\":[{\"key\":\"image\",\"selector\":\".image img\",\"attribute\":\"src\"}]}"

func TestUnmarshal(t *testing.T) {
	tt := []struct {
		name        string
		reqBody     string
		unmarshaler func(data []byte, v any) error
		resBody     *Body
		err         error
	}{
		{
			"handles Unmarshaler error",
			"",
			func(data []byte, v any) error {
				return fmt.Errorf("some error")
			},
			&Body{},
			fmt.Errorf("some error"),
		},
		{
			"handles successful Unmarshaler",
			htmlStub,
			json.Unmarshal,
			&Body{
				HTML: "<html></html>",
				Selectors: []Selector{{
					Key:       "image",
					Selector:  ".image img",
					Attribute: "src",
				}},
			},
			nil,
		},
		{
			"handles missing html field",
			"{\"selectors\":[{\"key\":\"name\",\"selector\":\".name\"}]}",
			json.Unmarshal,
			&Body{
				HTML: "",
				Selectors: []Selector{{
					Key:      "name",
					Selector: ".name",
				}},
			},
			fmt.Errorf("request body is missing 'html' field"),
		},
		{
			"handles missing selectors field",
			"{\"html\":\"<html></html>\"}",
			json.Unmarshal,
			&Body{
				HTML:      "<html></html>",
				Selectors: []Selector(nil),
			},
			fmt.Errorf("request body is missing 'selectors' field"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			unmarshaler = tc.unmarshaler
			resBody, err := Unmarshal(tc.reqBody)
			assert.Equal(t, tc.resBody, resBody)
			assert.Equal(t, tc.err, err)
		})
	}
}
