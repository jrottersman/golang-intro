package main

import "fmt"

func main() {
	half := func(x int) (int, bool) {
		var a int
		var b bool
		a = x / 2
		if x%2 == 0 {
			b = true
		} else {
			b = false
		}
		return a, b
	}
	y, z := half(0)
	fmt.Println(y)
	fmt.Println(z)
}
