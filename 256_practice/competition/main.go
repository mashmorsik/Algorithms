package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Participant struct {
	Time  int
	Place int
	ID    int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsNum int
	fmt.Fscanln(in, &setsNum)

	for i := 0; i < setsNum; i++ {
		var runnersNum int
		fmt.Fscanln(in, &runnersNum)

		timeResults, _ := in.ReadString('\n')

		parts := strings.Fields(timeResults)
		resultsArr := make([]int, len(parts))

		for j, part := range parts {
			fmt.Sscanf(part, "%d", &resultsArr[j])
		}

		participants := make([]Participant, runnersNum)

		for y := 0; y < runnersNum; y++ {
			participants[y] = Participant{Time: resultsArr[y], Place: y + 1, ID: y}
		}

		sort.Slice(participants, func(i, j int) bool {
			if participants[i].Time == participants[j].Time {
				return participants[i].Place < participants[j].Place
			}
			return participants[i].Time < participants[j].Time
		})

		currentPlace := 1
		for y := 0; y < runnersNum; y++ {
			participants[y].Place = currentPlace
			if y < runnersNum-1 {
				if y > 0 && (participants[y].Time != participants[y+1].Time &&
					participants[y-1].Time-participants[y].Time != 1 &&
					participants[y+1].Time-participants[y].Time != 1) {
					currentPlace = y + 2
				}
				if y == 0 && (participants[y+1].Time-participants[y].Time != 1 && participants[y].Time != participants[y+1].Time) {
					currentPlace = y + 2
				}
			}
		}

		sort.Slice(participants, func(i, j int) bool {
			return participants[i].ID < participants[j].ID
		})

		var results []int
		for _, participant := range participants {
			results = append(results, participant.Place)
		}

		strSlice := make([]string, len(results))
		for k, num := range results {
			strSlice[k] = strconv.Itoa(num)
		}

		result := strings.Join(strSlice, " ")

		fmt.Fprintln(out, result)
	}
}
