package y2022

import (
	"adventofcode/shared"
	"sort"
)

type monkey struct {
	items       []int
	operation   func(int) int
	target      func(int) int
	divisor     int
	inspections int
}

func Day11(input []string) (solution shared.Solution[int, int]) {
	solution.Part1 = solveDay11(getMonkeys(), 20, func(x int) int {
		return x / 3
	})
	monkeys := getMonkeys()
	divisor := 1
	for _, m := range monkeys {
		divisor *= m.divisor
	}
	solution.Part2 = solveDay11(monkeys, 10000, func(x int) int {
		return x % divisor
	})
	return
}

func solveDay11(monkeys map[int]*monkey, rounds int, decayFunc func(int) int) int {
	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]
			for _, item := range m.items {
				m.inspections += 1
				worryLevel := decayFunc(m.operation(item))
				next := monkeys[m.target(worryLevel%m.divisor)]
				next.items = append(next.items, worryLevel)
				m.items = m.items[1:]
			}
		}
	}
	insp := make([]int, len(monkeys))
	for _, monkey := range monkeys {
		insp = append(insp, monkey.inspections)
	}
	sort.Ints(insp)
	return insp[len(insp)-1] * insp[len(insp)-2]
}

func getMonkeys() map[int]*monkey {
	monkeys := map[int]*monkey{}
	monkeys[0] = &monkey{
		items:   []int{76, 88, 96, 97, 58, 61, 67},
		divisor: 3,
		operation: func(x int) int {
			return x * 19
		},
		target: func(rey2022der int) int {
			if rey2022der == 0 {
				return 2
			} else {
				return 3
			}
		}}
	monkeys[1] = &monkey{
		items:   []int{93, 71, 79, 83, 69, 70, 94, 98},
		divisor: 11,
		operation: func(x int) int {
			return x + 8
		},
		target: func(rey2022der int) int {
			if rey2022der == 0 {
				return 5
			} else {
				return 6
			}
		}}
	monkeys[2] = &monkey{
		items:   []int{50, 74, 67, 92, 61, 76},
		divisor: 19,
		operation: func(x int) int {
			return x * 13
		},
		target: func(rey2022der int) int {
			if rey2022der == 0 {
				return 3
			} else {
				return 1
			}
		}}
	monkeys[3] = &monkey{
		items:   []int{76, 92},
		divisor: 5,
		operation: func(x int) int {
			return x + 6
		},
		target: func(rey2022der int) int {
			if rey2022der == 0 {
				return 1
			} else {
				return 6
			}
		}}
	monkeys[4] = &monkey{
		items:   []int{74, 94, 55, 87, 62},
		divisor: 2,
		operation: func(x int) int {
			return x + 5
		},
		target: func(rey2022der int) int {
			if rey2022der == 0 {
				return 2
			} else {
				return 0
			}
		}}
	monkeys[5] = &monkey{
		items:   []int{59, 62, 53, 62},
		divisor: 7,
		operation: func(x int) int {
			return x * x
		},
		target: func(rey2022der int) int {
			if rey2022der == 0 {
				return 4
			} else {
				return 7
			}
		}}
	monkeys[6] = &monkey{
		items:   []int{62},
		divisor: 17,
		operation: func(x int) int {
			return x + 2
		},
		target: func(rey2022der int) int {
			if rey2022der == 0 {
				return 5
			} else {
				return 7
			}
		}}
	monkeys[7] = &monkey{
		items:   []int{85, 54, 53},
		divisor: 13,
		operation: func(x int) int {
			return x + 3
		},
		target: func(rey2022der int) int {
			if rey2022der == 0 {
				return 4
			} else {
				return 0
			}
		}}
	return monkeys
}
