package main

import "fmt"

func swap(x *int, y *int) (int, int) {
	tmp := *x
	*x = *y
	*y = tmp
	return *x, *y
}
func main() {
	x := 1
	y := 2
	a, b := swap(&x, &y)
	fmt.Println(a)
	fmt.Println(b)
}
