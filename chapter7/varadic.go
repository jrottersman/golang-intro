package main

import "fmt"

func greatest(args ...int) int {
	great := 0
	for _, v := range args {
		if v > great {
			great = v
		}
	}
	return great
}
func main() {
	fmt.Println(greatest(2, 4, 3))
}
