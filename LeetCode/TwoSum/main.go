package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 5, 5, 11}, 10))
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	for idx, _ := range nums {
		for i := idx + 1; i < len(nums)-1; i++ {
			if nums[idx]+nums[i] == target {
				return append(result, idx, i)
			}
		}
	}

	return result
}

func twoSumFaster(nums []int, target int) []int {
	// Create a map to store the indices of numbers
	numMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the complement needed to reach the target
		complement := target - num
		// Check if the complement exists in the map
		if idx, ok := numMap[complement]; ok {
			// If found, return the indices
			return []int{idx, i}
		}
		// Otherwise, add the current number and its index to the map
		numMap[num] = i
	}

	// If no solution is found, return an empty slice
	return []int{}
}
