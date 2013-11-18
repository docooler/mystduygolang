package main 

import "fmt"

func swap_string(x, y string) (string, string){
	return y, x
}

func main() {
	a, b := swap_string("love", "you")
	fmt.Println(a, b)
}