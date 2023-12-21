package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(GetSum(4, 100))
}

func GetSum(a, b int) int {
	arr := []int{a, b}
	sort.Ints(arr)
	var answer int

	for i := arr[0]; i < arr[1]+1; i++ {
		answer += i
	}
	return answer
}
