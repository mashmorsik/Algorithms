package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	bankA := make(map[string][]float32)
	bankB := make(map[string][]float32)
	bankC := make(map[string][]float32)

	banks := []map[string][]float32{bankA, bankB, bankC}

	var setsNum int
	fmt.Fscanln(in, &setsNum)

	for i := 0; i < setsNum; i++ {
		result := make([]float32, 0)

		for _, currBank := range banks {
			for j := 0; j < 6; j++ {
				line, _ := in.ReadString('\n')
				var arr []int
				for _, str := range strings.Fields(line) {
					num, err := strconv.Atoi(str)
					if err != nil {
						continue
					}
					arr = append(arr, num)
				}
				switch {
				case j < 2:
					currBank["RUB"] = append(currBank["RUB"], float32(arr[1])/float32(arr[0]))
				case j == 2 || j == 3:
					currBank["USD"] = append(currBank["USD"], float32(arr[1])/float32(arr[0]))
				default:
					currBank["EUR"] = append(currBank["EUR"], float32(arr[1])/float32(arr[0]))
				}
			}
		}

		for _, bank := range banks {
			result = append(result, bank["RUB"][0])
		}

		// RUS -> EUR

		rate1 := bankA["RUB"][1] * bankB["EUR"][1]
		result = append(result, rate1)

		rate2 := bankB["RUB"][1] * bankC["EUR"][1]
		result = append(result, rate2)

		rate3 := bankC["RUB"][1] * bankA["EUR"][1]
		result = append(result, rate3)

		rate4 := bankA["RUB"][1] * bankC["EUR"][1]
		result = append(result, rate4)

		rate5 := bankB["RUB"][1] * bankA["EUR"][1]
		result = append(result, rate5)

		rate6 := bankC["RUB"][1] * bankB["EUR"][1]
		result = append(result, rate6)

		//

		slices.Sort(result)
		fmt.Fprintln(out, result[len(result)-1])
	}
}
