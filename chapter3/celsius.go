package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	celsius := 1.8*input + 32

	fmt.Println(celsius)
}
