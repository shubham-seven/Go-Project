package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (sq square) area() float64 {
	return sq.length * sq.length
}
func (c circle) area() float64 {
	return (c.radius * c.radius) * (math.Pi)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		length: 10,
	}
	c := circle{
		radius: 12.345,
	}

	info(c)
	info(sq)

}
