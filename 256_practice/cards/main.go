package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numberOfFriendsAndCards string
	var cards string

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	numberOfFriendsAndCards, _ = in.ReadString('\n')
	cards, _ = in.ReadString('\n')

	numberArr := strings.Fields(numberOfFriendsAndCards)
	cardArr := strings.Fields(cards)

	cardsGiven, err := strconv.Atoi(numberArr[1])
	if err != nil {
		log.Fatal(err)
	}

	cardsMap := make(map[int]bool)

	for i := 1; i <= cardsGiven; i++ {
		cardsMap[i] = true
	}

	cardsFriendsArr := make([]int, len(cardArr))

	for i, str := range cardArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		cardsFriendsArr[i] = num
	}

	maxAttempts := cardsGiven
	for idx := range cardsFriendsArr {
		attempts := 0
		desired := cardsFriendsArr[idx] + 1
		for cardsMap[desired] == false {
			desired++
			attempts++
			if attempts > maxAttempts {
				fmt.Fprintln(out, -1)
				return
			}
		}
		cardsMap[desired] = false
		cardsFriendsArr[idx] = desired
	}

	result := make([]string, len(cardsFriendsArr))
	for i, num := range cardsFriendsArr {
		result[i] = strconv.Itoa(num)
	}

	fmt.Fprintln(out, strings.Join(result, " "))
}
