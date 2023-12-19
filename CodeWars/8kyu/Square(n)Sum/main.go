package main

import "fmt"

func main() {
	fmt.Println(SquareSum([]int{2, 2, 2}))
}

func SquareSum(numbers []int) int {
	var answer int
	for _, number := range numbers {
		number = number * number
		answer += number
	}
	return answer
}
