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

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscanln(in, &s)

		isPossible := true

		letters := []rune(s)
		firstLetter := letters[0]
		allSame := true
		if firstLetter != letters[len(letters)-1] {
			fmt.Fprintln(out, "NO")
			continue
		}

		for _, letter := range letters {
			if letter != firstLetter {
				allSame = false
				break
			}
		}

		if allSame {
			fmt.Fprintln(out, "YES")
			continue
		}

		var count int
		for _, letter := range letters {
			if letter == firstLetter {
				count++
			}
		}

		if count <= len(letters)/2 {
			isPossible = false
		}

		if count > len(letters)/2 {
			for y := 1; y < len(letters)-1; y++ {
				if letters[y] != firstLetter && (letters[y-1] != firstLetter || letters[y+1] != firstLetter) {
					isPossible = false
					break
				}
			}
		}

		if isPossible {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
