package main

import "fmt"

func main() {
	fmt.Println(returnSet([]int{3, 5, 1, 8, 1}, []int{1, 5, 1, 9, 0}))
}

func returnSet(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem] += 1
		}
	}
	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}
	return result
}
