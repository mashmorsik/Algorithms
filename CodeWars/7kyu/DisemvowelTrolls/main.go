package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Disemvowel("This is trolling"))
}

func Disemvowel(comment string) string {
	for _, vowel := range "aiueoAIUEO" {
		comment = strings.ReplaceAll(comment, string(vowel), "")
	}
	return comment
}
