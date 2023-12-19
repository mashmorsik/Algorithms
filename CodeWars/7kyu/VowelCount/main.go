package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GetCount("Awesome"))
}

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "o", "u", "i"}
	letters := strings.Split(str, "")

	for _, letter := range letters {
		for _, vowel := range vowels {
			if letter == vowel {
				count++
			}
		}
	}

	return count
}
