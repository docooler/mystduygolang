package main 

import "fmt"

func fibonacci() func() int {
	return func() int{
		return 1
	}
}


func  main() {
	f := fibonacci()
	for i:=0; i<10; i++{
		fmt.Println(f())
	}
}