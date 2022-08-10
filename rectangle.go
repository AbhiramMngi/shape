package Square

type Rectangle struct {
	length, width float64
}

func NewRectangle(length, width float64) Rectangle {
	if length <= 0 || width <= 0 {
		panic("length or width should be positive")
	}
	return Rectangle{length, width}
}

func (rectangle Rectangle) FindArea() float64 {
	if rectangle.length == 1.0 && rectangle.width == 1.0 {
		return 1.0
	}
	if rectangle.length == 1 {
		return rectangle.width
	}
	if rectangle.width == 1 {
		return rectangle.length
	}
	return rectangle.length * rectangle.width
}

func (rectangle Rectangle) FindPerimeter() float64 {
	if rectangle.length == 1 && rectangle.width == 1 {
		return 4.0
	}
	if rectangle.length == 1 {
		return 2 + 2*rectangle.width
	}
	if rectangle.width == 1 {
		return 2 + 2*rectangle.length
	}
	return (rectangle.length + rectangle.width) * 2
}
