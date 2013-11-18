package main 

import (
		"fmt"
		"math"
		)

type Rec struct {
	with, height float64
}

type Circle struct {
	radius float64
}

func (r Rec) area() float64 {
	return r.with * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func  main() {
	r1 := Rec{12, 2}
	r2 := Rec{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is : ", r1.area() )
	fmt.Println("Area of R2 is : ", r2.area() )
	fmt.Println("Area of c1 is : ", c1.area() )
	fmt.Println("Area of c2 is : ", c2.area() )
}