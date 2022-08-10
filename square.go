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

func (square Square) FindArea() float64 {
	if square.side == 1.0 {
		return 1.0
	}
	return square.side * square.side
}

func (square Square) FindPerimeter() float64 {
	if square.side == 1.0 {
		return 4.0
	}
	return square.side * 4
}
