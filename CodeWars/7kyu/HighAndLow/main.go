package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}

func HighAndLow(numbers string) string {
	numbersArr := strings.Split(numbers, " ")
	var intArr []int

	for _, number := range numbersArr {
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println(err)
		}
		intArr = append(intArr, numberInt)
	}

	sort.Ints(intArr)
	var high = strconv.Itoa(intArr[len(intArr)-1])
	var low = strconv.Itoa(intArr[0])
	return high + " " + low
}
