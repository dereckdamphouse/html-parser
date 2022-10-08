package res

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testsCleanUp() {
	os.Setenv("ALLOWORIGIN", "")
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

	testsCleanUp()
}
