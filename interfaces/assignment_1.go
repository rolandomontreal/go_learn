// Assignment for interfaces

package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type triangle struct {
	base float64
	height float64
}

// func main() {
// 	s := square{ side: 2.5 }
// 	printArea(s)
// }

func (s square) area() float64 {
	return s.side * s.side
}

func (t triangle) area() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println("Area: ", s.area())
}