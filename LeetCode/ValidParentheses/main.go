package main

import "fmt"

func main() {
	fmt.Println(isValid("{()}()[{}]"))
}

func isValid(s string) bool {
	var stack []int32
	memory := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	if s != "" {
		for _, char := range s {
			if len(stack) > 0 && stack[len(stack)-1] == memory[char] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, char)
			}
		}
	}

	return len(stack) == 0
}
