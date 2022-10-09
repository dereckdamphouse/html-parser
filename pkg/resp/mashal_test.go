package resp

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	tt := []struct {
		name      string
		found     any
		marshaler func(v any) ([]byte, error)
		res       string
		err       error
	}{
		{
			"handles Marshal error",
			"",
			func(v any) ([]byte, error) {
				return []byte{}, fmt.Errorf("some error")
			},
			"",
			fmt.Errorf("some error"),
		},
		{
			"handles successful Marshal",
			"found properties object",
			json.Marshal,
			"{\"error\":\"\",\"found\":\"found properties object\"}",
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			marshaler = tc.marshaler
			res, err := Marshal(tc.found)
			assert.Equal(t, tc.res, res)
			assert.Equal(t, tc.err, err)
		})
	}
}
