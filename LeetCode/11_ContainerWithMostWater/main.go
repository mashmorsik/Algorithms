package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	var left, answer int
	right := len(height) - 1

	for right > left {
		newAns := intersection(height[left], height[right]) * (right - left)
		answer = int(math.Max(float64(answer), float64(newAns)))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return answer
}

func intersection(a, b int) int {
	if a < b {
		return a
	}
	return b
}
