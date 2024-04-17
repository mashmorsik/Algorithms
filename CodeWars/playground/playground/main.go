package main

import "fmt"

func main() {
	var x []int
	x = append(x, 0) // len 1, cap 1
	x = append(x, 1) // len 2, cap 2
	x = append(x, 2) // len 3, cap 4

	y := append(x, 3) // [0 1 2 3]
	z := append(x, 4) // [0 1 2 4]

	fmt.Println(y, z) // [0 1 2 4], [0 1 2 4]
}
