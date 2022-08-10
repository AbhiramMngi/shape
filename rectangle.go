package Square

type Rectangle struct {
	length, width float64
}

func NewRectangle(length, width float64) Rectangle {
	return Rectangle{length, width}
}

func (rectangle Rectangle) FindPerimeter() float64 {
	return 0.0
}

func (rectangle Rectangle) FindArea() float64 {
	return 0.0
}
