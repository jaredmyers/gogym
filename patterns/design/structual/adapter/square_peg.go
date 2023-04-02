// Adapter Pattern Example
package main

import (
	"log"
	"math"
)

type TestFit interface {
	getRadius() float64
}

type RoundHole struct {
	radius float64
}

func (r *RoundHole) getRadius() float64 {
	return r.radius
}

func (r *RoundHole) fits(peg TestFit) bool {
	return r.getRadius() >= peg.getRadius()
}

type RoundPeg struct {
	radius float64
}

func (r *RoundPeg) getRadius() float64 {
	return r.radius
}

type SquarePeg struct {
	width float64
}

func (s *SquarePeg) getWidth() float64 {
	return s.width
}

type SquarePegAdapter struct {
	squarePeg *SquarePeg
}

func (s *SquarePegAdapter) getRadius() float64 {
	return s.squarePeg.getWidth() * (math.Sqrt(2) / 2)
}

func main() {

	hole := &RoundHole{radius: 5}
	roundPeg := &RoundPeg{radius: 5}

	log.Println(hole.fits(roundPeg))

	smallSquarePeg := &SquarePeg{width: 5}
	largeSquarePeg := &SquarePeg{width: 10}

	smallSquareAdapter := &SquarePegAdapter{squarePeg: smallSquarePeg}
	largeSquareAdapter := &SquarePegAdapter{squarePeg: largeSquarePeg}

	log.Println(hole.fits(smallSquareAdapter))
	log.Println(hole.fits(largeSquareAdapter))
}
