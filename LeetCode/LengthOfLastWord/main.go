package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Happy Birthday "))
}

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	var noSpaces []string

	for _, w := range arr {
		if w != "" {
			noSpaces = append(noSpaces, w)
			fmt.Println(noSpaces)
		}
	}

	return len(noSpaces[len(noSpaces)-1])
}
