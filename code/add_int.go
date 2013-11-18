package main 

import "fmt"

func add (in_x int, in_y int) int {
	return in_x + in_y
}
func add1 (in_x, in_y int) int {
	return in_y + in_x
}
func main() {
	fmt.Println(add(42, 43))
	fmt.Println(add(42, 43))
}