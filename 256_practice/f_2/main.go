package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var lettersNum, testsNum int
	fmt.Fscanln(in, &lettersNum, &testsNum)

	var letters string
	letters, _ = in.ReadString('\n')

	counts := make(map[rune]int)

	for _, letter := range letters {
		if letter != ' ' && letter != '\n' {
			counts[letter]++
		}
	}

	for i := 0; i < testsNum; i++ {
		var test string
		fmt.Fscanln(in, &test)

		countsCopy := make(map[rune]int)
		for k, v := range counts {
			countsCopy[k] = v
		}

		allLettersFound := true
		for _, ch := range test {
			if count, exists := countsCopy[ch]; exists && count > 0 {
				countsCopy[ch]--
			} else {
				allLettersFound = false
				break
			}
		}

		for _, count := range countsCopy {
			if count > 0 {
				allLettersFound = false
				break
			}
		}

		if allLettersFound {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
