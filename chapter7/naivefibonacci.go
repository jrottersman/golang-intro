package main

import "fmt"

func fib(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		x := fib(x-1) + fib(x-2)
		return x
	}
}
func main() {
	fmt.Println(fib(8))
	fmt.Println(fib(7))
	fmt.Println(fib(5))
}
