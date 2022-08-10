package Square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	t.Run("should check if the rectangle object is created", func(t *testing.T) {
		assert.IsType(t, Rectangle{}, NewRectangle(1.0, 2.0))
	})
	t.Run("should check for panic when length or width is not positive", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(-20.0, 32.4)
		})
	})
	t.Run("should check for no panic when length and width are positive", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewRectangle(2.0, 3.0)
		})
	})
}

func TestFindRectangleArea(t *testing.T) {
	t.Run("should check the area of the rectangle when the length and width are 1", func(t *testing.T) {
		assert.Equal(t, 1.0, NewRectangle(1.0, 1.0).FindArea())
	})
	t.Run("should check the area of the rectangle when the length is 1 and width is not equal to 1", func(t *testing.T) {
		assert.Equal(t, 5.0, NewRectangle(1.0, 5.0).FindArea())
	})
	t.Run("should check the area of the rectangle when the length is not equal to 1 and width is 1", func(t *testing.T) {
		assert.Equal(t, 5.0, NewRectangle(5.0, 1.0).FindArea())
	})
	t.Run("should check the area of the rectangle when the length is not equal to 1 and width is not equal to 1", func(t *testing.T) {
		assert.Equal(t, 27.0, NewRectangle(9.0, 3.0).FindArea())
	})
	t.Run("should check the area of the rectangle when the length and width are swapped", func(t *testing.T) {
		assert.Equal(t, NewRectangle(3.0, 9.0).FindArea(), NewRectangle(9.0, 3.0).FindArea())
	})
}

func TestFindRectanglePerimeter(t *testing.T) {
	t.Run("should check the perimeter of the rectangle when length and width are 1", func(t *testing.T) {
		assert.Equal(t, 4.0, NewRectangle(1.0, 1.0).FindPerimeter())
	})
	t.Run("should check the perimeter of the rectangle when the length is 1 and width is not equal to 1", func(t *testing.T) {
		assert.Equal(t, 22.0, NewRectangle(1.0, 10.0).FindPerimeter())
	})
	t.Run("should check the perimeter of the rectangle when the length is not equal to 1 and width is 1", func(t *testing.T) {
		assert.Equal(t, 22.0, NewRectangle(10.0, 1.0).FindPerimeter())
	})
	t.Run("should check the perimeter of the rectangle when the length is not equal to 1 and width is not equal to 1", func(t *testing.T) {
		assert.Equal(t, 44.0, NewRectangle(10.0, 12.0).FindPerimeter())
	})
	t.Run("should check the perimeter of the rectangle when the length and width are swapped", func(t *testing.T) {
		assert.Equal(t, NewRectangle(18.0, 19.0).FindPerimeter(), NewRectangle(19.0, 18.0).FindPerimeter())
	})
}
