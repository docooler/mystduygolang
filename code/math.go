package main 

import (
		"fmt"
		"math"
		)

func  main() {
	fmt.Println("now you have %g problems.",
		math.Nextafter(2,3))
	fmt.Println("I forget told you Pi is :", math.Pi)
}