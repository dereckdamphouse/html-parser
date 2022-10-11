package resp

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dereckdamphouse/html-parser/pkg/html"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	tt := []struct {
		name      string
		parsed    html.Parsed
		marshaler func(v any) ([]byte, error)
		res       string
		err       error
	}{
		{
			"handles Marshal error",
			make(html.Parsed),
			func(v any) ([]byte, error) {
				return []byte{}, fmt.Errorf("some error")
			},
			"",
			fmt.Errorf("some error"),
		},
		{
			"handles successful Marshal",
			html.Parsed{
				"title": {"title1"},
				"image": {"imageUrl1, imageUrl2"},
			},
			json.Marshal,
			"{\"image\":[\"imageUrl1, imageUrl2\"],\"title\":[\"title1\"]}",
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			marshaler = tc.marshaler
			res, err := Marshal(tc.parsed)
			assert.Equal(t, tc.res, res)
			assert.Equal(t, tc.err, err)
		})
	}
}
