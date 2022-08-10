package Square

type Square struct {
	side float64
}

func NewSquare(side float64) Square {
	if side <= 0 {
		panic("square side should be positive")
	}
	return Square{side}
}

func (square Square) FindPerimeter() float64 {
	return 0.0
}

func (square Square) FindArea() float64 {
	return 0.0
}
