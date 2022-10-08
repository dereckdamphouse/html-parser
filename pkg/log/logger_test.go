package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	t.Run("logger has init", func(t *testing.T) {
		assert.NotEqual(t, &zap.Logger{}, Instance)
	})
}
