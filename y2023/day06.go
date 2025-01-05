package y2023

import (
	"adventofcode/shared"
	"strconv"
	"strings"
)

func Day06(input []string) (solution shared.Solution[int, int]) {
	distances := shared.IntSlice(strings.Fields(input[1])[1:])
	solution.Part1 = 1
	for i, time := range shared.IntSlice(strings.Fields(input[0])[1:]) {
		tmp := 0
		for v := 1; v < time; v++ {
			tmp += shared.BoolToInt((time-v)*v > distances[i])
		}
		solution.Part1 *= tmp
	}
	time, _ := strconv.Atoi(strings.Join(strings.Fields(input[0])[1:], ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(input[1])[1:], ""))
	for v := 1; v < time; v++ {
		solution.Part2 += shared.BoolToInt((time-v)*v > distance)
	}
	return
}
