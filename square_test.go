package Square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSquare(t *testing.T) {
	t.Run("should check if the square object is created", func(t *testing.T) {
		assert.IsType(t, Square{}, NewSquare(5.0))
	})
	t.Run("should check for panic when a negative side is given", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-9.0)
		})
	})
}
