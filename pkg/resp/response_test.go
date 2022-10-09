package resp

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	res := Error(http.StatusBadRequest, "some error")
	assert.Equal(t, res.Body, fmt.Sprintf("{\"error\":\"%s\",\"found\":{}}", "some error"))
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

func TestSuccess(t *testing.T) {
	res := Success("{\"found\":{}}")
	assert.Equal(t, res.Body, "{\"found\":{}}")
	assert.Equal(t, res.StatusCode, 200)
}

func TestAllowOrigin(t *testing.T) {
	tt := []struct {
		name        string
		ao          string
		aoEnvSetter func()
	}{
		{
			"handles empty env",
			"",
			func() {},
		},
		{
			"handles short env",
			"",
			func() {
				os.Setenv("ALLOWORIGIN", "''")
			},
		},
		{
			"handles wildcard env",
			"*",
			func() {
				os.Setenv("ALLOWORIGIN", "'*'")
			},
		},
		{
			"handles URL env",
			"https://www.google.com",
			func() {
				os.Setenv("ALLOWORIGIN", "'https://www.google.com'")
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.aoEnvSetter()
			assert.Equal(t, tc.ao, allowOrigin())
		})
	}

	os.Setenv("ALLOWORIGIN", "")
}
