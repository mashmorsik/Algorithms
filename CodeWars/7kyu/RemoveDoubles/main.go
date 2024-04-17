package main

import (
	"fmt"
	"slices"
)

func main() {
	//fmt.Println(RemoveSame([]int{5, -7, 1, 1, 1, 6, -7, 6}))
	fmt.Println(RemoveDoubles([]int{5, -7, 1, 1, 1, 6, -7, 6}))
	fmt.Println(Remove([]int{5, -7, 1, 1, 1, 6, -7, 6}))
}

func RemoveDoubles(numbers []int) []int {
	ind := 0

	slices.Sort(numbers)
	y := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if y != numbers[i] {
			numbers[ind] = y
			ind++
		}
		y = numbers[i]
	}

	numbers[ind] = numbers[len(numbers)-1]
	ind++

	return numbers[:ind]
}

func Remove(nums []int) []int {
	unique := make(map[int]struct{}, len(nums))
	ind := 0

	for _, n := range nums {
		if _, ok := unique[n]; !ok {
			unique[n] = struct{}{}
			nums[ind] = n
			ind++
		}
	}

	return nums[:ind]
}

//func RemoveSame(nums []int) []int {
//	var result []int
//	unique := make(map[int]struct{})
//
//	for _, n := range nums {
//		unique[n] = struct{}{}
//	}
//
//	for n := range unique {
//		result = append(result, n)
//	}
//
//	return result
//}

//func Remove(nums []int) []int {
//	unique := make(map[int]struct{})
//	resultIdx := 0
//
//	for _, n := range nums {
//		if _, ok := unique[n]; !ok {
//			unique[n] = struct{}{}
//			nums[resultIdx] = n
//			resultIdx++
//		}
//	}
//
//	return nums[:resultIdx]
//}
