package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func newSquare(x int, y int, length int) (*Square, error) {
	if length < 1 {
		return nil, fmt.Errorf("%d is not a valid length", length)
	}

	center := Point{
		X: x,
		Y: y,
	}

	square := &Square{
		Center: center,
		Length: length,
	}

	return square, nil
}

func main() {
	_, err := newSquare(1, 2, -1)
	if err != nil {
		fmt.Println(err)
	}

	square, err := newSquare(2, 2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", square)
	}

	fmt.Printf("Area: %d\n", square.Area())

	square.Move(4, 5)

	fmt.Printf("Moved square: %+v\n", square)
}
