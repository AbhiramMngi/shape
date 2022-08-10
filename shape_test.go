package Square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShape(t *testing.T) {
	t.Run("should check if the shape interface is implemented by square", func(t *testing.T) {
		assert.Implements(t, new(Shape), NewSquare(5.0))
	})
	t.Run("should check if the shape interface is implemented by rectangle", func(t *testing.T) {
		assert.Implements(t, new(Shape), NewRectangle(5.0, 6.0))
	})
}
