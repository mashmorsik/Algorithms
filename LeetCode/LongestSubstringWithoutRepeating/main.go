package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	memory := make(map[rune]struct{})
	var left, right, res int

	strRunes := []rune(s)

	for right < len(strRunes) {
		c := strRunes[right]
		if _, ok := memory[c]; !ok {
			memory[c] = struct{}{}
			res = int(math.Max(float64(res), float64(right-left+1)))
			right++
		} else {
			delete(memory, strRunes[left])
			left++
		}
	}

	return res
}
