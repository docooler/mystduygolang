package main 

import "fmt"

func loop_while(max int) int{
     sum := 1
     for sum < max {
     	sum += sum
     }
     return sum
}
func main() {
	sum := 0
	for i:=0; i<100; i++{
		sum += i
	}
	fmt.Println("sum is :%g", sum)

	j :=99
	sum = 0
	for ; j>=0; j-- {
		sum +=j
	}
	fmt.Println(sum)
	fmt.Println(loop_while(1000))
}