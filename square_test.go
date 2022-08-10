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
	t.Run("should check for no panic when a positive side is given", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewSquare(9.0)
		})
	})
}

func TestFindSquareArea(t *testing.T) {
	t.Run("should check if the area of the square is correct when side is 1", func(t *testing.T) {
		assert.Equal(t, 1.0, NewSquare(1.0).FindArea())
	})
	t.Run("should check if the area of the square is correct", func(t *testing.T) {
		assert.Equal(t, 100.0, NewSquare(10.0).FindArea())
	})
}

func TestFindSquarePerimeter(t *testing.T) {
	t.Run("should check if the perimeter of the square when the side is 1", func(t *testing.T) {
		assert.Equal(t, 4.0, NewSquare(1.0).FindPerimeter())
	})
	t.Run("should check if the perimeter of the square", func(t *testing.T) {
		assert.Equal(t, 20.0, NewSquare(5.0).FindPerimeter())
	})
}
