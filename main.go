package main

import (
	"fmt"
)

func main() {
	var x []int = []int{10, 5, 4, -10, 1000}
	var y []int = []int{0, 0}

	copy(y, x[3:5])
	z := append(y, -999)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
