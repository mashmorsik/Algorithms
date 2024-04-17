package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(containsDuplicate([]int{0, 1, 6, 7, -8, 2, 2, 0, 1, -8, 9, 9}))
	fmt.Println(containsDuplicate([]int{0, 1, 6, 7}))
}

func containsDuplicate(nums []int) bool {
	slices.Sort(nums)
	x := nums[0]
	isDouble := false

	for i := 1; i < len(nums); i++ {
		if x == nums[i] {
			isDouble = true
		}
		x = nums[i]
	}
	return isDouble
}
