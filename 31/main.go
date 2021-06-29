package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func distanceBeetwenPoints(p1 Point, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

var (
	p1 = NewPoint(3, 4)
	p2 = NewPoint(6, 8)
)

func main() {
	fmt.Printf("Point 1: %v\nPoint 2: %v\nDistance: %v\n", p1, p2, distanceBeetwenPoints(p1, p2))
}
