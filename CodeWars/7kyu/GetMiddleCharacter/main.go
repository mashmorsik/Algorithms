package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetMiddle("Hello"))
}

func GetMiddle(word string) string {
	if len(word) > 2 {
		return GetMiddle(word[1 : len(word)-1])
	}

	return word
}
